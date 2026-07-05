import glfw
from PIL import Image
from OpenGL.GL import *
import numpy as np
import math
import time
import ctypes

angle_x, angle_y = 0.0, 0.0
wireframe = False
scale = 0.8
texture_enabled = True

pos = np.array([-0.3, 0.0, 0.0], dtype=np.float32)
vel = np.array([0.5, 0.4, 0.3], dtype=np.float32)
BOUND = 1.1
last_time = time.time()

def three_point_projection():
    return [0.87, -0.09, 0.98, 0.49,
            0.0,   0.98, 0.35, 0.17,
            0.5,   0.15,-1.7, -0.85,
            0.0,   0.0,  1.0,  2.0]

vertex_shader_source = """
#version 330 core
in vec3 position;
in vec3 vertexNormal;
in vec2 texcoord;

uniform mat4 model;
uniform mat4 view;
uniform mat4 projection;

out vec3 fragPos;
out vec3 fragNormal;
out vec2 Texcoord;

void main() {
    vec4 worldPos = model * vec4(position,1.0);
    fragPos = worldPos.xyz;
    fragNormal = mat3(transpose(inverse(model))) * vertexNormal;

    Texcoord = texcoord;
    gl_Position = projection * view * model * vec4(position, 1.0);
}
"""

fragment_shader_source = """
#version 330 core
in vec3 fragPos;
in vec3 fragNormal;
in vec2 Texcoord;

uniform sampler2D texSampler;
uniform bool textureEnabled;

uniform vec3 lightPos;
uniform vec3 lightDir;
uniform vec3 lightAmbient;
uniform vec3 lightDiffuse;
uniform vec3 lightSpecular;
uniform float spotCutoff;
uniform float spotExponent;

uniform vec3 materialAmbient;
uniform vec3 materialDiffuse;
uniform vec3 materialSpecular;
uniform float materialShininess;
uniform vec3 globalAmbient;

out vec4 outColor;

void main() {
    vec3 N = normalize(fragNormal);
    vec3 L = normalize(lightPos - fragPos);
    vec3 V = normalize(-fragPos);
    vec3 R = reflect(-L, N);

    vec3 spotDir = normalize(lightDir);
    float cosTheta = dot(-L, spotDir);
    float spotFactor = cosTheta < spotCutoff ? 0.0 : pow(cosTheta, spotExponent);

    vec3 ambient = globalAmbient * materialAmbient + lightAmbient * materialAmbient * spotFactor;
    float diff = max(dot(N, L), 0.0);
    vec3 diffuse = diff * materialDiffuse * lightDiffuse * spotFactor;
    float spec = pow(max(dot(R, V), 0.0), materialShininess);
    vec3 specular = spec * materialSpecular * lightSpecular * spotFactor;

    vec3 lightingColor = ambient + diffuse + specular;
    vec4 texColor = texture(texSampler, Texcoord);
    vec3 finalColor = textureEnabled ? texColor.rgb * lightingColor : lightingColor;
    outColor = vec4(finalColor, 1.0);
}
"""

def build_sphere(radius=1.0, segments=10, sides=20):
    vertices = []
    indices = []
    for i in range(segments):
        theta1 = -math.pi/2 + i/segments * math.pi
        theta2 = -math.pi/2 + (i+1)/segments * math.pi
        y1, y2 = radius*math.sin(theta1), radius*math.sin(theta2)
        r1, r2 = radius*math.cos(theta1), radius*math.cos(theta2)
        for j in range(sides+1):
            phi = j * 2*math.pi / sides
            x1, z1 = r1 * math.cos(phi), r1 * math.sin(phi)
            x2, z2 = r2 * math.cos(phi), r2 * math.sin(phi)
            nx1, ny1, nz1 = x1/radius, y1/radius, z1/radius
            nx2, ny2, nz2 = x2/radius, y2/radius, z2/radius
            u, v1 = phi/(2*math.pi), (y1/radius+1)/2
            v2 = (y2/radius+1)/2
            vertices.extend([x1,y1,z1, nx1,ny1,nz1, u,v1])
            vertices.extend([x2,y2,z2, nx2,ny2,nz2, u,v2])
        base = i*2*(sides+1)
        for j in range(sides):
            i0, i1, i2, i3 = base+2*j, base+2*j+1, base+2*j+2, base+2*j+3
            indices.extend([i0,i2,i1, i1,i2,i3])
    return np.array(vertices, dtype=np.float32), np.array(indices, dtype=np.uint32)

def create_texture(path='3.jpg'):
    img = Image.open(path).convert('RGB').resize((128,128))
    w, h = img.size
    tex = glGenTextures(1)
    glBindTexture(GL_TEXTURE_2D, tex)
    glTexImage2D(GL_TEXTURE_2D, 0, GL_RGB, w, h, 0, GL_RGB, GL_UNSIGNED_BYTE, img.tobytes())
    glTexParameteri(GL_TEXTURE_2D, GL_TEXTURE_MIN_FILTER, GL_LINEAR)
    glTexParameteri(GL_TEXTURE_2D, GL_TEXTURE_MAG_FILTER, GL_LINEAR)
    glTexParameteri(GL_TEXTURE_2D, GL_TEXTURE_WRAP_S, GL_REPEAT)
    glTexParameteri(GL_TEXTURE_2D, GL_TEXTURE_WRAP_T, GL_REPEAT)
    return tex

def translate_matrix(v):
    m = np.eye(4, dtype=np.float32)
    m[:3,3] = v
    return m

def scale_matrix(s):
    return np.diag([s,s,s,1]).astype(np.float32)

def rotate_x_matrix(deg):
    a = math.radians(deg)
    c, s = math.cos(a), math.sin(a)
    return np.array([[1,0,0,0],[0,c,-s,0],[0,s,c,0],[0,0,0,1]], dtype=np.float32)

def rotate_y_matrix(deg):
    a = math.radians(deg)
    c, s = math.cos(a), math.sin(a)
    return np.array([[c,0,s,0],[0,1,0,0],[-s,0,c,0],[0,0,0,1]], dtype=np.float32)

def compile_shader(src, stype, name):
    sh = glCreateShader(stype)
    glShaderSource(sh, src)
    glCompileShader(sh)
    return sh

def make_program(vsrc, fsrc):
    vertex = compile_shader(vsrc, GL_VERTEX_SHADER, "vertex")
    fragment = compile_shader(fsrc, GL_FRAGMENT_SHADER, "fragment")
    prog = glCreateProgram()
    glAttachShader(prog, vertex)
    glAttachShader(prog, fragment)
    glLinkProgram(prog)

    glDetachShader(prog, vertex); 
    glDetachShader(prog, fragment)
    glDeleteShader(vertex); 
    glDeleteShader(fragment)
    glValidateProgram(prog)
    return prog

def key_callback(window, key, scancode, action, mods):
    global angle_x, angle_y, wireframe, scale, texture_enabled
    if action in (glfw.PRESS, glfw.REPEAT):
        if key == glfw.KEY_RIGHT:    angle_y -= 3
        elif key == glfw.KEY_LEFT:   angle_y += 3
        elif key == glfw.KEY_UP:     angle_x -= 3
        elif key == glfw.KEY_DOWN:   angle_x += 3
        elif key in (glfw.KEY_KP_ADD, glfw.KEY_EQUAL): scale += 0.1
        elif key in (glfw.KEY_KP_SUBTRACT, glfw.KEY_MINUS): scale = max(0.1, scale-0.1)
        elif key == glfw.KEY_M: wireframe = not wireframe
        elif key == glfw.KEY_T: texture_enabled = not texture_enabled

def display(shader_program, vao, num_indices, texture):
    global angle_x, angle_y, scale, pos, vel, last_time, texture_enabled

    current_time = time.time()
    dt = current_time - last_time
    last_time = current_time
    if dt > 0.1: dt = 0.016

    for i in range(3):
        pos[i] += vel[i] * dt
        if pos[i] > BOUND:
            pos[i] = BOUND
            vel[i] = -vel[i]
        elif pos[i] < -BOUND:
            pos[i] = -BOUND   
            vel[i] = -vel[i]

    glClear(GL_COLOR_BUFFER_BIT | GL_DEPTH_BUFFER_BIT)

    model = translate_matrix(pos) @ scale_matrix(scale) @ rotate_x_matrix(angle_x) @ rotate_y_matrix(angle_y)
    view = np.eye(4, dtype=np.float32)
    projection = three_point_projection()

    glUseProgram(shader_program)

    glUniformMatrix4fv(glGetUniformLocation(shader_program, 'model'), 1, GL_TRUE, model)
    glUniformMatrix4fv(glGetUniformLocation(shader_program, 'view'), 1, GL_TRUE, view)
    glUniformMatrix4fv(glGetUniformLocation(shader_program, 'projection'), 1, GL_FALSE, projection)

    glUniform3fv(glGetUniformLocation(shader_program, 'lightPos'), 1,  [0, 3, 0])
    glUniform3fv(glGetUniformLocation(shader_program, 'lightDir'), 1,  [0, -1, 0])
    glUniform3fv(glGetUniformLocation(shader_program, 'lightAmbient'), 1, [0.3, 0.3, 0.3])
    glUniform3fv(glGetUniformLocation(shader_program, 'lightDiffuse'), 1, [3, 3, 3])
    glUniform3fv(glGetUniformLocation(shader_program, 'lightSpecular'), 1,[3, 3, 3])
    glUniform1f(glGetUniformLocation(shader_program, 'spotCutoff'), math.cos(math.radians(70)))
    glUniform1f(glGetUniformLocation(shader_program, 'spotExponent'), 70)

    glUniform3fv(glGetUniformLocation(shader_program, 'materialAmbient'), 1, [1, 1, 1])
    glUniform3fv(glGetUniformLocation(shader_program, 'materialDiffuse'), 1, [1, 1, 1])
    glUniform3fv(glGetUniformLocation(shader_program, 'materialSpecular'), 1, [0.8, 0.8, 0.8])
    glUniform1f(glGetUniformLocation(shader_program, 'materialShininess'), 64)
    glUniform3fv(glGetUniformLocation(shader_program, 'globalAmbient'), 1, [0.3, 0.3, 0.3])
    glUniform1i(glGetUniformLocation(shader_program, 'textureEnabled'), 1 if texture_enabled else 0)
    glUniform1i(glGetUniformLocation(shader_program, 'texSampler'), 0)

    if texture_enabled:
        glActiveTexture(GL_TEXTURE0)
        glBindTexture(GL_TEXTURE_2D, texture)
    else:
        glBindTexture(GL_TEXTURE_2D, 0)

    glPolygonMode(GL_FRONT_AND_BACK, GL_LINE if wireframe else GL_FILL)
    glBindVertexArray(vao)
    glDrawElements(GL_TRIANGLES, num_indices, GL_UNSIGNED_INT, None)

def make_program(vsrc, fsrc):
    vertex = compile_shader(vsrc, GL_VERTEX_SHADER, "vertex")
    fragment = compile_shader(fsrc, GL_FRAGMENT_SHADER, "fragment")
    prog = glCreateProgram()
    glAttachShader(prog, vertex)
    glAttachShader(prog, fragment)
    glLinkProgram(prog)
    return prog

def main():
    global last_time
    if not glfw.init():
        return

    glfw.window_hint(glfw.CONTEXT_VERSION_MAJOR, 3)
    glfw.window_hint(glfw.CONTEXT_VERSION_MINOR, 3)
    glfw.window_hint(glfw.OPENGL_PROFILE, glfw.OPENGL_CORE_PROFILE)
    glfw.window_hint(glfw.OPENGL_FORWARD_COMPAT, GL_TRUE)
    glfw.window_hint(glfw.COCOA_RETINA_FRAMEBUFFER, GL_FALSE)

    window = glfw.create_window(800, 800, "lab8", None, None)
    if not window:
        glfw.terminate()
        return

    glfw.make_context_current(window)
    glfw.set_key_callback(window, key_callback)

    glEnable(GL_DEPTH_TEST)
    glClearColor(0.15, 0.15, 0.15, 1.0)

    temp_vao = glGenVertexArrays(1)
    glBindVertexArray(temp_vao)

    shader_program = make_program(vertex_shader_source, fragment_shader_source)

    glBindVertexArray(0)
    glDeleteVertexArrays(1, [temp_vao])

    sphere_vertices, sphere_indices = build_sphere()

    vao = glGenVertexArrays(1)
    vbo, ebo = glGenBuffers(2)

    glBindVertexArray(vao)
    glBindBuffer(GL_ARRAY_BUFFER, vbo)
    glBufferData(GL_ARRAY_BUFFER, sphere_vertices.nbytes, sphere_vertices, GL_STATIC_DRAW)
    glBindBuffer(GL_ELEMENT_ARRAY_BUFFER, ebo)
    glBufferData(GL_ELEMENT_ARRAY_BUFFER, sphere_indices.nbytes, sphere_indices, GL_STATIC_DRAW)

    stride = 8 * np.dtype(np.float32).itemsize

    pos_loc = glGetAttribLocation(shader_program, "position")
    glEnableVertexAttribArray(pos_loc)
    glVertexAttribPointer(pos_loc, 3, GL_FLOAT, GL_FALSE, stride, ctypes.c_void_p(0))

    norm_loc   = glGetAttribLocation(shader_program, "vertexNormal")
    glEnableVertexAttribArray(norm_loc)
    glVertexAttribPointer(norm_loc, 3, GL_FLOAT, GL_FALSE, stride, ctypes.c_void_p(12))

    texcoord_loc = glGetAttribLocation(shader_program, "texcoord")
    glEnableVertexAttribArray(texcoord_loc)
    glVertexAttribPointer(texcoord_loc, 2, GL_FLOAT, GL_FALSE, stride, ctypes.c_void_p(24))

    texture = create_texture('3.jpg')
    last_time = time.time()

    while not glfw.window_should_close(window):
        display(shader_program, vao, len(sphere_indices), texture)
        glfw.swap_buffers(window)
        glfw.poll_events()

    glDeleteVertexArrays(1, [vao])
    glDeleteBuffers(1, [vbo])
    glDeleteBuffers(1, [ebo])
    glDeleteProgram(shader_program)
    glDeleteTextures([texture])

    glfw.destroy_window(window)
    glfw.terminate()

if __name__ == "__main__":
    main()