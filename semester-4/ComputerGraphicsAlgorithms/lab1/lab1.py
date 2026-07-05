import glfw
from OpenGL.GL import *

delta = 0.2
angle = 0.0
posx = 0.0
posy = 0.0
color = 0.0
magic_num = 0.0
flag = True

def main():
    if not glfw.init():
        return
    window = glfw.create_window(640, 640, "Lab1", None, None)
    if not window:
        glfw.terminate()
        return
    glfw.make_context_current(window)
    glfw.set_key_callback(window, key_callback)
    glfw.set_scroll_callback(window, scroll_callback)
    while not glfw.window_should_close(window):
        display(window)
    glfw.destroy_window(window)
    glfw.terminate()


def display(window):
    global angle
    global magic_num
    global color
    glClear(GL_COLOR_BUFFER_BIT)
    glLoadIdentity()
    glClearColor(1.0, 1.0, 1.0, 1.0)
    glPushMatrix()
    glRotatef(angle, 0, 0, 1)
    glBegin(GL_POLYGON)
    glColor3f(0.0 + color,0.95 + color,0.6 + color)
    glVertex2f(posx + 0.5,posy + 0.5)
    glColor3f(0.96 + color,0.6 + color,0.0 + color)
    glVertex2f(posx + -0.5,posy + 0.5)
    glColor3f(0.32 + color,0.1 + color,0.71 + color)
    glVertex2f(posx + 0.25 - magic_num,posy + 0.25 - magic_num)
    glColor3f(0.13 + color,0.32 + color,0.65 + color)
    glVertex2f(posx + 0.5,posy + -0.5)
    glEnd()

    glBegin(GL_POLYGON)
    glColor3f(0.4 + color,0.6 + color,0.6 + color)
    glVertex2f(posx - 0.25 + magic_num,posy - 0.25 + magic_num)
    glColor3f(0. + color,0.0 + color,0.89 + color)
    glVertex2f(posx + -0.5,posy + 0.5)
    glColor3f(0.45 + color,0.66 + color,0.49 + color)
    glVertex2f(posx - 0.5,posy - 0.5)
    glColor3f(0.78 + color,0.23 + color, 0.89 + color)
    glVertex2f(posx + 0.5,posy + -0.5)
    glEnd()

    glPopMatrix()
    angle += delta
    glfw.swap_buffers(window)
    glfw.poll_events()

def key_callback(window, key, scancode, action, mods):
    global magic_num
    global flag
    if action == glfw.PRESS:
        if key == glfw.KEY_RIGHT:
            if flag:
                magic_num += 0.03
        if key == 263: # glfw.KEY_LEFT
            if flag:
                magic_num += -0.03

def scroll_callback(window, xoffset, yoffset):
    global color
    if (xoffset > 0):
        color -= yoffset/100
    else:
        color += yoffset/100


if __name__ == "__main__":
    main()