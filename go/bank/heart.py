5# import math
# from turtle import *
# def hearta(k):
#   shapea = 15 * math.sin(k) ** 3
#   return shapea

# def heartb(k):
#   shapeb = 12*math.cos(k)-5*\
#   math.cos(2*k)-2*\
#   math.cos(3*k)-\
#   math.cos(4*k)
#   return shapeb

# speed(9000)
# bgcolor("white")

# for i in range(6000):
#   goto(hearta(i)*20, heartb(i)*20)
#   for j in range(5):
#     color("red")
#     goto(0,0)
# done()
import time
from threading import Thread, Lock
import sys

lock = Lock()

def animate_text(text, delay=0.1):
    with lock:
        for char in text:
            sys.stdout.write(char)
            sys.stdout.flush()
            time.sleep(delay)
        print()

def sing_lyric(lyric, delay, speed):
    time.sleep(delay)
    animate_text(lyric, speed)

def sing_song():
    lyrics = [
        ("And I reach for your hand, but it's cold, you pull away again", 0.07),
        ("And I wonder what's on your mind", 0.07),
        ("And then you say to me you made a dumb mistake", 0.07),
        ("You start to tremble and your voice begins to break", 0.07),
        ("You say the cigarettes on the counter weren't your friend's", 0.07),
        ("They were my mate's", 0.07),
        ("And I feel the color draining from my face", 0.07),
        ("And my friend said", 0.07),
        ("I know you love her, but it's over, mate", 0.08),
        ("It doesn't matter, put the phone away", 0.08),
        ("It's never easy to walk away", 0.08),
        ("Let her go, it'll be alright", 0.08)
    ]
    delays = [ 6.5, 9.0, 12.0, 15.0, 18.0, 20.0, 22.0, 24.0, 27.0, 30.0, 33.0, 36.0]
    
    threads = []
    for i in range(len(lyrics)):
        lyric, speed = lyrics[i]
        t = Thread(target=sing_lyric, args=(lyric, delays[i], speed))
        threads.append(t)
        t.start()
    for thread in threads:
        thread.join()

if __name__ == "__main__":
    sing_song()
