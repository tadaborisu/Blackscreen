import pygame
from win32api import GetSystemMetrics

pygame.init()
pygame.display.set_caption('Blackscreen')
gameIcon = pygame.image.load('icon.ico')
pygame.display.set_icon(gameIcon)
screen = pygame.display.set_mode((GetSystemMetrics(0), GetSystemMetrics(1)), pygame.FULLSCREEN)

running = True
while running:

    for event in pygame.event.get():
        if event.type == pygame.QUIT:
            running = False

    screen.fill((0, 0, 0))