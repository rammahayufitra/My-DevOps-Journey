from flask import Flask, Response, request, send_from_directory, abort
from waitress import serve
from inference import Q, single_thread
import queue
from threading import Thread


app = Flask("__name__")

@app.route("/ai_image", methods=["GET"])
def get_single_image():
    return Response(Q().get_single_frame(q), mimetype='multipart/x-mixed-replace; boundary=frame')

if __name__ == '__main__':
    q = queue.Queue()
    frame_queue = 3
    thread = Thread(target=single_thread, args=(q,frame_queue), daemon=True).start()
    serve(app,host="0.0.0.0", port=5000)
