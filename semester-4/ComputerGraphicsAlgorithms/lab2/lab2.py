import glfw
from OpenGL.GL import *

angle_x, angle_y = 0, 0
wireframe = False
scale = 1.0

def three_point_projection():
    M_3P = [
            0.87, -0.09, 0.98, 0.49,
            0, 0.98, 0.35, 0.17,
            0.5, 0.15, -1.7, -0.85, 
            0, 0, 1, 2,
        ]
    return M_3P

def main():
    if not glfw.init():
        return
    window = glfw.create_window(800, 800, "lab2", None, None)
    if not window:
        glfw.terminate()
        return
    glfw.make_context_current(window)
    glfw.set_key_callback(window, key_callback)
    
    glEnable(GL_DEPTH_TEST)
    glClearColor(0.2, 0.2, 0.2, 1.0)

    while not glfw.window_should_close(window):
        display()
        glfw.swap_buffers(window)
        glfw.poll_events()
    
    glfw.destroy_window(window)
    glfw.terminate()

def draw_cube():
    face_colors = [
        [0.8, 0.0, 0.0],
        [0.0, 0.8, 0.0],
        [0.0, 0.0, 0.8],
        [0.8, 0.8, 0.0],
        [0.8, 0.0, 0.8],
        [0.0, 0.8, 0.8],
    ]

    vertices = [
        [-0.3, -0.3, -0.3], [0.3, -0.3, -0.3], [0.3, 0.3, -0.3], [-0.3, 0.3, -0.3],
        [-0.3, -0.3, 0.3], [0.3, -0.3, 0.3], [0.3, 0.3, 0.3], [-0.3, 0.3, 0.3]
    ]
    faces = [
        [0, 1, 2, 3], [4, 5, 6, 7], [0, 4, 7, 3],
        [1, 5, 6, 2], [3, 2, 6, 7], [0, 1, 5, 4]
    ]
    
    glPolygonMode(GL_FRONT_AND_BACK, GL_LINE if wireframe else GL_FILL)
    
    for i, face in enumerate(faces):
        glBegin(GL_QUADS)
        glColor3fv(face_colors[i])
        for vertex in face:
            glVertex3fv(vertices[vertex])
        glEnd()

def display():
    global angle_x, angle_y, scale
    
    glLoadIdentity()
    glClear(GL_COLOR_BUFFER_BIT)
    glClear(GL_DEPTH_BUFFER_BIT)
    glMatrixMode(GL_PROJECTION)
    glPushMatrix()

    glTranslatef(0.6, 0.6, 0)
    glMultMatrixf(three_point_projection())
    draw_cube()
    glPopMatrix()
    
    glTranslatef(-0.3, 0, 0)
    glScalef(scale, scale, scale)
    glMultMatrixf(three_point_projection())
    glRotatef(angle_x, 1, 0, 0)
    glRotatef(angle_y, 0, 1, 0)
    draw_cube()


def key_callback(window, key, scancode, action, mods):
    global angle_x, angle_y, wireframe, scale
    if action in [glfw.PRESS, glfw.REPEAT]:
        if key == glfw.KEY_RIGHT:
            angle_y -= 3
        if key == glfw.KEY_LEFT:
            angle_y += 3
        if key == glfw.KEY_UP:
            angle_x -= 3
        if key == glfw.KEY_DOWN:
            angle_x += 3
        if key == glfw.KEY_KP_ADD or key == glfw.KEY_EQUAL:
            scale += 0.1
        if key == glfw.KEY_KP_SUBTRACT or key == glfw.KEY_MINUS:
            scale = max(0.1, scale - 0.1)
        if key == glfw.KEY_M:
            wireframe = not wireframe

if __name__ == "__main__":
    main()