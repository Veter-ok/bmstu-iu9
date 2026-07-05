import glfw
from PIL import Image
from OpenGL.GL import *
import math
import time
import ctypes
import ctypes.util

class PerformanceTimer:
    def __init__(self):
        self.libc = ctypes.CDLL(ctypes.util.find_library('c'))
        try:
            self._get_time_ns = self.libc.clock_gettime_nsec_np
            self._get_time_ns.argtypes = [ctypes.c_int]
            self._get_time_ns.restype = ctypes.c_uint64
            self.use_mach = False
        except AttributeError:
            self.libkernel = ctypes.CDLL(ctypes.util.find_library('kernel'))
            self._mach_absolute_time = self.libkernel.mach_absolute_time
            self._mach_absolute_time.argtypes = []
            self._mach_absolute_time.restype = ctypes.c_uint64
            class mach_timebase_info(ctypes.Structure):
                _fields_ = [("numer", ctypes.c_uint32), ("denom", ctypes.c_uint32)]
            self._timebase = mach_timebase_info()
            self.libkernel.mach_timebase_info(ctypes.byref(self._timebase))
            self.use_mach = True

    def get_time_ns(self):
        if self.use_mach:
            ticks = self._mach_absolute_time()
            return (ticks * self._timebase.numer) // self._timebase.denom
        else:
            return self._get_time_ns(0)

angle_x, angle_y = 0, 0
wireframe = False
scale = 0.8
texture_enabled = True
texture_id = None

pos = [-0.3, 0.0, 0.0]
vel = [0.5, 0.4, 0.3]
BOUND = 1.1
SPHERE_RADIUS = 1.0
last_time = time.time()
timer = PerformanceTimer()
sphere_display_list = None


def three_point_projection():
    return [
        0.87, -0.09, 0.98, 0.49,
        0,     0.98, 0.35, 0.17,
        0.5,   0.15,-1.7, -0.85,
        0,     0,    1,    2
    ]

def setup_lighting():
    glEnable(GL_LIGHTING)
    glEnable(GL_LIGHT0)

    glLightfv(GL_LIGHT0, GL_AMBIENT, [0.3, 0.3, 0.3, 1])
    glLightfv(GL_LIGHT0, GL_DIFFUSE, [3, 3, 3, 1])
    glLightfv(GL_LIGHT0, GL_SPECULAR, [3, 3, 3, 1])
    glLightfv(GL_LIGHT0, GL_POSITION, [0, 3, 0, 0])

    glLightf(GL_LIGHT0, GL_SPOT_CUTOFF, 70)
    glLightfv(GL_LIGHT0, GL_SPOT_DIRECTION, [0, -1, 0])
    glLightf(GL_LIGHT0, GL_SPOT_EXPONENT, 70)

    glLightModelfv(GL_LIGHT_MODEL_AMBIENT,  [0.3, 0.3, 0.3, 1.0])
    glLightModeli(GL_LIGHT_MODEL_LOCAL_VIEWER, GL_FALSE)
    glLightModeli(GL_LIGHT_MODEL_TWO_SIDE, GL_FALSE)

def setup_material():
    glEnable(GL_COLOR_MATERIAL)
    glColorMaterial(GL_FRONT, GL_AMBIENT_AND_DIFFUSE)
    glMaterialfv(GL_FRONT, GL_SPECULAR, [0.8, 0.8, 0.8, 1.0])
    glMaterialf(GL_FRONT, GL_SHININESS, 64.0)

def create_checkerboard_texture():
    img = Image.open('3.jpg').resize((128, 128))
    img_data = img.tobytes()
    width, height = img.size
    texture = glGenTextures(1)
    glBindTexture(GL_TEXTURE_2D, texture)
    glTexImage2D(GL_TEXTURE_2D, 0, GL_RGB, width, height, 0, GL_RGB, GL_UNSIGNED_BYTE, img_data)
    glTexParameteri(GL_TEXTURE_2D, GL_TEXTURE_MIN_FILTER, GL_LINEAR)
    glTexParameteri(GL_TEXTURE_2D, GL_TEXTURE_MAG_FILTER, GL_LINEAR)
    glTexParameteri(GL_TEXTURE_2D, GL_TEXTURE_WRAP_S, GL_REPEAT)
    glTexParameteri(GL_TEXTURE_2D, GL_TEXTURE_WRAP_T, GL_REPEAT)
    return texture

def generate_sphere_vertex_arrays(segments=10, sides=20, R=SPHERE_RADIUS):
    slices = []
    for i in range(segments):
        theta1 = -math.pi/2 + (i / segments) * math.pi
        theta2 = -math.pi/2 + ((i + 1) / segments) * math.pi
        y1 = R * math.sin(theta1)
        y2 = R * math.sin(theta2)
        r1 = R * math.cos(theta1)
        r2 = R * math.cos(theta2)
        height = y2 - y1
        if r1 >= 0 and r2 >= 0:
            vertices = []
            normals = []
            texcoords = []
            angle_step = 2 * math.pi / sides
            angle = 0.0
            for _ in range(sides + 1):
                x1 = r1 * math.cos(angle)
                z1 = r1 * math.sin(angle)
                x2 = r2 * math.cos(angle)
                z2 = r2 * math.sin(angle)
                y_top = y1 + height

                if r1 > 0.001:
                    nx1, ny1, nz1 = x1/R, y1/R, z1/R
                else:
                    nx1, ny1, nz1 = 0.0, -1.0, 0.0
                if r2 > 0.001:
                    nx2, ny2, nz2 = x2/R, y_top/R, z2/R
                else:
                    nx2, ny2, nz2 = 0.0, 1.0, 0.0

                u = angle / (2 * math.pi)
                v1 = (y1 / R + 1.0) / 2.0
                v2 = (y_top / R + 1.0) / 2.0

                vertices.extend([x1, y1, z1])
                normals.extend([nx1, ny1, nz1])
                texcoords.extend([u, v1])

                vertices.extend([x2, y_top, z2])
                normals.extend([nx2, ny2, nz2])
                texcoords.extend([u, v2])

                angle += angle_step

            vertex_count = len(vertices) // 3
            vert_arr = (ctypes.c_float * len(vertices))(*vertices)
            norm_arr = (ctypes.c_float * len(normals))(*normals)
            tex_arr = (ctypes.c_float * len(texcoords))(*texcoords)
            slices.append((vert_arr, norm_arr, tex_arr, vertex_count))
    return slices

def init_display_list():
    global sphere_display_list
    slices = generate_sphere_vertex_arrays()
    sphere_display_list = glGenLists(1)
    glNewList(sphere_display_list, GL_COMPILE)
    glEnableClientState(GL_VERTEX_ARRAY)
    glEnableClientState(GL_NORMAL_ARRAY)
    glEnableClientState(GL_TEXTURE_COORD_ARRAY)
    for vert_arr, norm_arr, tex_arr, vcount in slices:
        glVertexPointer(3, GL_FLOAT, 0, vert_arr)
        glNormalPointer(GL_FLOAT, 0, norm_arr)
        glTexCoordPointer(2, GL_FLOAT, 0, tex_arr)
        glDrawArrays(GL_TRIANGLE_STRIP, 0, vcount)
    glDisableClientState(GL_VERTEX_ARRAY)
    glDisableClientState(GL_NORMAL_ARRAY)
    glDisableClientState(GL_TEXTURE_COORD_ARRAY)
    glEndList()

def display():
    global angle_x, angle_y, scale, pos, vel, last_time, texture_enabled, texture_id
    global sphere_display_list

    current_time = time.time()
    dt = current_time - last_time
    last_time = current_time
    if dt > 0.1:
        dt = 0.016
    for i in range(3):
        pos[i] += vel[i] * dt
        if pos[i] > BOUND:
            pos[i] = BOUND
            vel[i] = -vel[i]
        elif pos[i] < -BOUND:
            pos[i] = -BOUND
            vel[i] = -vel[i]

    glClear(GL_COLOR_BUFFER_BIT | GL_DEPTH_BUFFER_BIT)

    glMatrixMode(GL_PROJECTION)
    glLoadIdentity()
    glMultMatrixf(three_point_projection())
    glMatrixMode(GL_MODELVIEW)
    glLoadIdentity()

    if texture_enabled:
        glEnable(GL_TEXTURE_2D)
        glBindTexture(GL_TEXTURE_2D, texture_id)
        glTexEnvi(GL_TEXTURE_ENV, GL_TEXTURE_ENV_MODE, GL_MODULATE)
    else:
        glDisable(GL_TEXTURE_2D)

    glPolygonMode(GL_FRONT_AND_BACK, GL_LINE if wireframe else GL_FILL)

    glTranslatef(pos[0], pos[1], pos[2])
    glScalef(scale, scale, scale)
    glRotatef(angle_x, 1, 0, 0)
    glRotatef(angle_y, 0, 1, 0)
    glColor3f(1, 1, 1)

    start = timer.get_time_ns()
    if sphere_display_list is not None:
        glCallList(sphere_display_list)
    end = timer.get_time_ns()
    print(f"Frame time: {(end - start) / 1e6:.5f} ms")

    glDisable(GL_TEXTURE_2D)

def key_callback(window, key, scancode, action, mods):
    global angle_x, angle_y, wireframe, scale, texture_enabled
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
        if key == glfw.KEY_T:
            texture_enabled = not texture_enabled

def main():
    global texture_id, last_time, sphere_display_list

    if not glfw.init():
        return
    window = glfw.create_window(800, 800, "lab7", None, None)
    if not window:
        glfw.terminate()
        return
    glfw.make_context_current(window)
    glfw.set_key_callback(window, key_callback)

    glEnable(GL_DEPTH_TEST)
    glClearColor(0.15, 0.15, 0.15, 1.0)

    glEnable(GL_CULL_FACE)
    glCullFace(GL_BACK)
    glFrontFace(GL_CCW)

    setup_lighting()
    setup_material()
    texture_id = create_checkerboard_texture()

    init_display_list()

    last_time = time.time()
    while not glfw.window_should_close(window):
        display()
        glfw.swap_buffers(window)
        glfw.poll_events()

    if sphere_display_list:
        glDeleteLists(sphere_display_list, 1)
    glfw.destroy_window(window)
    glfw.terminate()

if __name__ == "__main__":
    main()