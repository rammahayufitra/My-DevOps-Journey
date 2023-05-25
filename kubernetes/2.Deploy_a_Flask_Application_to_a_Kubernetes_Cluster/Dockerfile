FROM python:3.8

RUN apt-get update 
RUN apt-get upgrade -y


WORKDIR /app

COPY . .

RUN pip3 install --upgrade pip
RUN pip3 install flask 
RUN pip3 install opencv-python==4.1.2.30
RUN pip3 install waitress 

EXPOSE 5000

CMD ["python3", "/app/interface.py"]