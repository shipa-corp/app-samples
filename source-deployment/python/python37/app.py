from flask import Flask
import sys

app = Flask(__name__)

@app.route("/")
def hello():
    return "{0}.{1}".format(sys.version_info.major, sys.version_info.minor)

@app.route("/healthcheck")
def healthcheck():
    return "OKEY"

if __name__ == "__main__":
    app.run()
