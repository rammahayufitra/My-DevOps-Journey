import cv2 
import time
import queue
from queue import Empty 

class Q:
    """ Implement Queue (FIFO)
        Main methods:
        PUT - put frames overlay AI inside QUEUE
        CLEAR - clear frames in QUEUE if frames > the maximum number of frame in QUEUE
        GET - get frames from QUEUE 
    """
    def __init__(self):
        pass
    def clearQueue(self, q):
        """ clear frames in queue """ 
        try:
            while not q.empty():
                q.get_nowait()
        except Empty:
            pass
    def get_single_frame(self, q):
        """ get frame"""
        while True:
            if not q.empty():
                frame = q.get()
                resized = cv2.resize(frame, (1280, 720), interpolation=cv2.INTER_AREA)
                frame = cv2.imencode('.jpg', resized)[1].tobytes()
                yield (b'--frame\r\nContent-Type: image/jpeg\r\n\r\n' + frame + b'\r\n\r\n')
                q.task_done()
            else:
                time.sleep(0.1)
    def put_frame(self, frame, q, frame_queue):
        """ Put frame and clear frame"""
        try:
            q.put(frame)
        except:
            pass
        if q.qsize() >= frame_queue:
            self.clearQueue(q)

def single_thread(q, frame_queue):
    FILENAME = "http://192.168.200.229:8080/video"
    cam = cv2.VideoCapture(FILENAME) 
    while True:
        ret, frame = cam.read()
        Q().put_frame(frame, q, frame_queue)


