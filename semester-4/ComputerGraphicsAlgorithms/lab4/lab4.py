import glfw
from OpenGL.GL import *
import numpy as np

WINDOW_WIDTH = 800
WINDOW_HEIGHT = 600

class Polygon():
    def __init__(self):
        self.points = []
        self.edges_color = (255, 0, 0)
        self.fill_color = (76, 145, 65)
        self.buffer = np.full((WINDOW_HEIGHT, WINDOW_WIDTH, 3), (255, 255, 255), dtype=np.uint8)

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
        x0, y0 = point1
        x1, y1 = point2
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

    def fill_with_zatravka(self, x0, y0):
        stack = [(x0, y0)]
        while len(stack) != 0:
            x, y = stack.pop()
            x_left, x_right = x, x
            while not self.is_edge_point(x_left, y):
                self.set_pixel(x_left, y, self.fill_color)
                x_left -= 1
            while not self.is_edge_point(x_right, y):
                self.set_pixel(x_right, y, self.fill_color)
                x_right += 1
            x_left += 1
            x_right -= 1
            for xi in range(x_left, x_right + 1):   
                if self.is_free(xi, y - 1, x_left):
                    stack.append((xi, y - 1))
                if self.is_free(xi, y + 1, x_left):
                    stack.append((xi, y + 1))  

    def weighted_filter_3x3(self):
        kernel = np.array([[1, 2, 1],
                            [2, 4, 2],
                            [1, 2, 1]], dtype=np.float32)

        filtered = np.copy(self.buffer)
        h, w, _ = self.buffer.shape
        for y in range(1, h - 1):
            for x in range(1, w - 1):
                region = self.buffer[y - 1:y + 2, x - 1:x + 2]
                for c in range(3):
                    value = np.sum(region[:, :, c] * kernel) / kernel.sum()
                    filtered[y, x, c] = np.clip(value, 0, 255)
        self.buffer = filtered
    
polygons = Polygon()

def main():
    if not glfw.init():
        return
    window = glfw.create_window(WINDOW_WIDTH, WINDOW_HEIGHT, "Lab4", None, None)
    if not window:
        glfw.terminate()
        return
    glfw.make_context_current(window)
    glfw.set_key_callback(window, key_callback)
    glfw.set_mouse_button_callback(window, mouse_button_callback)
    glfw.set_framebuffer_size_callback(window, framebuffer_size_callback)
    
    glEnable(GL_TEXTURE_2D)
    texture = glGenTextures(1)
    while not glfw.window_should_close(window):
        glfw.poll_events()
        
        glBindTexture(GL_TEXTURE_2D, texture)
        glTexParameteri(GL_TEXTURE_2D, GL_TEXTURE_MIN_FILTER, GL_NEAREST)
        glTexParameteri(GL_TEXTURE_2D, GL_TEXTURE_MAG_FILTER, GL_NEAREST)
        glTexImage2D(GL_TEXTURE_2D, 0, GL_RGB, polygons.buffer.shape[1], polygons.buffer.shape[0],
                     0, GL_RGB, GL_UNSIGNED_BYTE, polygons.buffer)
        
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

def key_callback(window, key, scancode, action, mods):
    if action == glfw.PRESS:
        if key == glfw.KEY_F:
            polygons.weighted_filter_3x3()
        if key == glfw.KEY_C:
            polygons.clear_buffer()

def mouse_button_callback(window, button, action, mods):
    if action == glfw.PRESS:
        win_w, win_h = glfw.get_window_size(window)
        buf_w, buf_h = polygons.buffer.shape[1], polygons.buffer.shape[0]
        x, y = glfw.get_cursor_pos(window)
        x = int(x * buf_w / win_w)
        y = int(y * buf_h / win_h)
        y = buf_h - y
        if button == glfw.MOUSE_BUTTON_LEFT:
            polygons.add_point(int(x), int(y))
            if len(polygons.points) > 1:
                polygons.bresenham(polygons.points[-2], polygons.points[-1])
        if button == glfw.MOUSE_BUTTON_RIGHT:
            if len(polygons.points) > 2:
                polygons.bresenham(polygons.points[-1], polygons.points[0])
            polygons.fill_with_zatravka(int(x),int(y))

def framebuffer_size_callback(window, width, height):
    global WINDOW_WIDTH, WINDOW_HEIGHT
    WINDOW_WIDTH, WINDOW_HEIGHT = width, height
    glViewport(0, 0, width, height)

        
main()