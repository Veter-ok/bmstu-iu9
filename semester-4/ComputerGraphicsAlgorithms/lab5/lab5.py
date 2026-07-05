import glfw
from OpenGL.GL import *
import numpy as np

WINDOW_WIDTH = 800
WINDOW_HEIGHT = 600
IS_CLIP_POLYGON_FINISHED = False
EPS = 1e-6

class Polygon():
    def __init__(self, buffer, edges_color=(255, 0, 0), fill_color=(200, 200, 200)):
        self.points = []
        self.edges_color = edges_color
        self.fill_color = fill_color
        self.buffer = buffer

    def clear_buffer(self):
        self.buffer.fill(255)
        self.points = []

    def add_point(self, x, y):
        self.points.append((x, y))

    def set_pixel(self, x, y, color):
        if 0 <= x < self.buffer.shape[1] and 0 <= y < self.buffer.shape[0]:
            self.buffer[y, x] = color
    
    def is_edge_point(self, x, y):
        h, w = self.buffer.shape[:2]
        return not (0 <= y < h and 0 <= x < w) or tuple(self.buffer[y, x]) == self.edges_color

    def is_filled_point(self, x, y):
        h, w = self.buffer.shape[:2]
        return 0 <= y < h and 0 <= x < w and tuple(self.buffer[y, x]) == self.fill_color  

    def is_free(self, x, y, x_left):
        return not self.is_edge_point(x, y) and not self.is_filled_point(x, y) \
                   and (x == x_left or self.is_edge_point(x - 1, y) or self.is_filled_point(x - 1, y))

    def bresenham(self, point1, point2):
        x0, y0 = map(int, point1)
        x1, y1 = map(int, point2)
        dx = abs(x1 - x0)
        dy = abs(y1 - y0)
        sx = 1 if x0 < x1 else -1
        sy = 1 if y0 < y1 else -1
        
        if dx >= dy:
            e = -dx
            de = 2 * dy
            while True:
                self.set_pixel(x0, y0, self.edges_color)
                if x0 == x1:
                    break
                e += de
                x0 += sx
                if e >= 0:
                    y0 += sy
                    e -= 2 * dx
        else:
            e = -dy
            de = 2 * dx
            while True:
                self.set_pixel(x0, y0, self.edges_color)
                if y0 == y1:
                    break
                e += de
                y0 += sy
                if e >= 0:
                    x0 += sx
                    e -= 2 * dy

buffer = np.full((WINDOW_HEIGHT, WINDOW_WIDTH, 3), (255, 255, 255), dtype=np.uint8)
clip_polygon = Polygon(buffer, edges_color=(255, 0, 0))
clipping_polygon = Polygon(buffer, edges_color=(0, 0, 255), fill_color=(200, 200, 255))

def main():
    if not glfw.init():
        return
    window = glfw.create_window(WINDOW_WIDTH, WINDOW_HEIGHT, "Lab5", None, None)
    if not window:
        glfw.terminate()
        return
    glfw.make_context_current(window)
    glfw.set_key_callback(window, key_callback)
    glfw.set_mouse_button_callback(window, mouse_button_callback)
    
    glEnable(GL_TEXTURE_2D)
    texture = glGenTextures(1)
    while not glfw.window_should_close(window):
        glfw.poll_events()
        
        glBindTexture(GL_TEXTURE_2D, texture)
        glTexParameteri(GL_TEXTURE_2D, GL_TEXTURE_MIN_FILTER, GL_NEAREST)
        glTexParameteri(GL_TEXTURE_2D, GL_TEXTURE_MAG_FILTER, GL_NEAREST)
        glTexImage2D(GL_TEXTURE_2D, 0, GL_RGB, buffer.shape[1], buffer.shape[0],
                     0, GL_RGB, GL_UNSIGNED_BYTE, buffer)
        
        glClear(GL_COLOR_BUFFER_BIT)
        glClearColor(1.0, 1.0, 1.0, 1.0)
        glLoadIdentity()
        
        glBegin(GL_QUADS)
        glTexCoord2f(0, 0); glVertex2f(-1, -1)
        glTexCoord2f(1, 0); glVertex2f(1, -1)
        glTexCoord2f(1, 1); glVertex2f(1, 1)
        glTexCoord2f(0, 1); glVertex2f(-1, 1)
        glEnd()
        
        glfw.swap_buffers(window)
    glfw.destroy_window(window)
    glfw.terminate()

def orientation(a, b, c):
    return (b[0] - a[0]) * (c[1] - a[1]) - (b[1] - a[1]) * (c[0] - a[0])

def inside(p, a, b):
    return orientation(a, b, p) >= 0

def intersection(a1, a2, b1, b2):
    x1, y1 = a1
    x2, y2 = a2
    x3, y3 = b1
    x4, y4 = b2
    d = (x1-x2)*(y3-y4) - (y1-y2)*(x3-x4)
    if d == 0:
        return None
    px = ((x1*y2 - y1*x2)*(x3-x4) - (x1-x2)*(x3*y4 - y3*x4)) / d
    py = ((x1*y2 - y1*x2)*(y3-y4) - (y1-y2)*(x3*y4 - y3*x4)) / d
    return (px, py)

def Sutherland_Hodgman(object, clip):
    output = object
    for i in range(len(clip)):
        input_list = output
        output = []
        if not input_list:
            return []
        A = clip[i - 1]
        B = clip[i]
        S = input_list[-1]
        for P in input_list:
            if inside(P, A, B):
                if not inside(S, A, B):
                    I = intersection(S, P, A, B)
                    if I:
                        output.append(I)
                output.append(P)
            elif inside(S, A, B):
                I = intersection(S, P, A, B)
                if I:
                    output.append(I)
            S = P
    return output

def key_callback(window, key, scancode, action, mods):
    global IS_CLIP_POLYGON_FINISHED
    if action == glfw.PRESS:
        if key == glfw.KEY_C:
            IS_CLIP_POLYGON_FINISHED = False
            clip_polygon.clear_buffer()
            clipping_polygon.clear_buffer()
        if key == glfw.KEY_V:
            buffer.fill(255)
            result = Sutherland_Hodgman(clipping_polygon.points, clip_polygon.points)
            clipping_polygon.points = result
            for i in range(len(clip_polygon.points)):
                p1 = clip_polygon.points[i-1]
                p2 = clip_polygon.points[i]
                clip_polygon.bresenham(p1, p2)
            for i in range(len(clipping_polygon.points)):
                p1 = clipping_polygon.points[i-1]
                p2 = clipping_polygon.points[i]
                clipping_polygon.bresenham(p1, p2)
        if key == glfw.KEY_SPACE:
            if IS_CLIP_POLYGON_FINISHED:
                if len(clipping_polygon.points) > 2:
                    clipping_polygon.bresenham(clipping_polygon.points[-1], clipping_polygon.points[0])
            else:
                IS_CLIP_POLYGON_FINISHED = True
                if len(clip_polygon.points) > 2:
                    clip_polygon.bresenham(clip_polygon.points[-1], clip_polygon.points[0])

def mouse_button_callback(window, button, action, mods):
    if action == glfw.PRESS:
        win_w, win_h = glfw.get_window_size(window)
        buf_w, buf_h = buffer.shape[1], buffer.shape[0]
        x, y = glfw.get_cursor_pos(window)
        x = int(x * buf_w / win_w)
        y = int(y * buf_h / win_h)
        y = buf_h - y
        if button == glfw.MOUSE_BUTTON_LEFT:
            if IS_CLIP_POLYGON_FINISHED:
                clipping_polygon.add_point(int(x), int(y))
                if len(clipping_polygon.points) > 1:
                    clipping_polygon.bresenham(clipping_polygon.points[-2], clipping_polygon.points[-1])
            else:
                clip_polygon.add_point(int(x), int(y))
                if len(clip_polygon.points) > 1:
                    clip_polygon.bresenham(clip_polygon.points[-2], clip_polygon.points[-1])

        
main()