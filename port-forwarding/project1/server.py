from flask import Flask, send_from_directory, abort 

app = Flask(__name__)

app.config["CLIENT_IMAGES"] = '/home/ramma/My-DevOps-Journey/port-forwarding/project1/'

@app.route("/")
def home():
    return ({"HAY":"ramma"})

@app.route("/get_image/<image_name>")
def get_image(image_name):
    try:
        print(app.config["CLIENT_IMAGES"]) 
        return send_from_directory(app.config["CLIENT_IMAGES"], filename=image_name, as_attachment=False)
    except FileNotFoundError:
        abort(404)


app.run()
