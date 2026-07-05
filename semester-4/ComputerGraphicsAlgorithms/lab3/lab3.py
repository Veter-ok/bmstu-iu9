import glfw
from OpenGL.GL import *
import math

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
    window = glfw.create_window(800, 800, "lab3", None, None)
    if not window:
        glfw.terminate()
        return
    glfw.make_context_current(window)
    glfw.set_key_callback(window, key_callback)
    
    glEnable(GL_DEPTH_TEST)
    glClearColor(1, 1, 1, 1)

    while not glfw.window_should_close(window):
        display()
        glfw.swap_buffers(window)
        glfw.poll_events()
    
    glfw.destroy_window(window)
    glfw.terminate()

def draw_part(y, R1, R2, height, sides):
    angle_step = 2*math.pi / sides
    angle = 0
    glBegin(GL_QUAD_STRIP)
    for _ in range(sides+1):
        glColor3f(0.5, 0.1, 0.1)
        x1 = R1 * math.cos(angle)
        z1 = R1 * math.sin(angle)
        x2 = R2 * math.cos(angle)
        z2 = R2 * math.sin(angle)
        glVertex3f(x1, y, z1)
        glVertex3f(x2, y+height, z2)
        angle += angle_step
    glEnd()
    

def draw_sphere():
    R = 1
    segments = 10
    glPolygonMode(GL_FRONT_AND_BACK, GL_LINE if wireframe else GL_FILL)
    
    for i in range(segments):
        theta1 = -math.pi/2 + (i / segments) * math.pi
        theta2 = -math.pi/2 + ((i + 1) / segments) * math.pi
        
        y1 = R * math.sin(theta1)
        y2 = R * math.sin(theta2)
        r1 = R * math.cos(theta1)
        r2 = R * math.cos(theta2)
        
        height = y2 - y1
        if r1 > 0 and r2 > 0:
            draw_part(y1, r1, r2, height, 20)


def display():
    global angle_x, angle_y, scale
    
    glLoadIdentity()
    glClear(GL_COLOR_BUFFER_BIT)
    glClear(GL_DEPTH_BUFFER_BIT)
    glMatrixMode(GL_PROJECTION)
    glPushMatrix()

    glTranslatef(0.6, 0.6, 0)
    glScalef(0.2, 0.2, 0.2)
    glMultMatrixf(three_point_projection())
    draw_sphere()
    glPopMatrix()
    
    glTranslatef(-0.3, 0, 0)
    glScalef(scale, scale, scale)
    glMultMatrixf(three_point_projection())
    glRotatef(angle_x, 1, 0, 0)
    glRotatef(angle_y, 0, 1, 0)
    draw_sphere()


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