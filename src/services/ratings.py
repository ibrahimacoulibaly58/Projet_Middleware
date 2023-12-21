import json
import requests
from flask_login import current_user

from src.models.http_exceptions import *
from src.schemas.rating import RatingSchema
import src.repositories.ratings as ratings_repository

ratings_url = "http://localhost:4000/ratings/"  


def get_rating(rating_id):
    response = requests.request(method="GET", url=ratings_url + rating_id)
    return response.json(), response.status_code


def get_all_ratings():
    response = requests.request(method="GET", url=ratings_url)
    return response.json(), response.status_code


def create_rating(rating_data):
    rating_schema = RatingSchema().loads(json.dumps(rating_data))
    response = requests.request(method="POST", url=ratings_url, json=rating_schema)
    if response.status_code != 201:
        return response.json(), response.status_code

    return response.json(), response.status_code


def modify_rating(rating_id, rating_data):
    response = requests.request(method="PUT", url=ratings_url + rating_id, json=rating_data)
    if response.status_code != 200:
        return response.json(), response.status_code

    return response.json(), response.status_code


def delete_rating(rating_id):
    response = requests.request(method="DELETE", url=ratings_url + rating_id)
    if response.status_code != 204:
        return response.json(), response.status_code

    return {}, 204


def get_rating_from_db(rating_id):
    return ratings_repository.get_rating(rating_id)


def rating_exists(rating_id):
    return get_rating_from_db(rating_id) is not None
