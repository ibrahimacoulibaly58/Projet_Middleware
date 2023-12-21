import json
import requests
from flask_login import current_user

from src.models.http_exceptions import *
from src.schemas.song import SongSchema
import src.repositories.songs as songs_repository

songs_url = "http://localhost:4000/songs/"  # Remplacez cela par l'URL réelle de votre API songs


def get_song(song_id):
    response = requests.request(method="GET", url=songs_url + song_id)
    return response.json(), response.status_code


def get_all_songs():
    response = requests.request(method="GET", url=songs_url)
    return response.json(), response.status_code


def create_song(song_data):
    song_schema = SongSchema().loads(json.dumps(song_data))
    response = requests.request(method="POST", url=songs_url, json=song_schema)
    if response.status_code != 201:
        return response.json(), response.status_code

    return response.json(), response.status_code


def modify_song(song_id, song_data):
    response = requests.request(method="PUT", url=songs_url + song_id, json=song_data)
    if response.status_code != 200:
        return response.json(), response.status_code

    return response.json(), response.status_code


def delete_song(song_id):
    response = requests.request(method="DELETE", url=songs_url + song_id)
    if response.status_code != 204:
        return response.json(), response.status_code

    return {}, 204


def get_song_from_db(song_id):
    return songs_repository.get_song(song_id)


def song_exists(song_id):
    return get_song_from_db(song_id) is not None
