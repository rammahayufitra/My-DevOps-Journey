from flask import Flask 

app = Flask(__name__)

@app.route("/data")
def home():
    return "welcom to deplot flask with nginx"

if __name__ == "__main__":
    app.run(debug=True)
