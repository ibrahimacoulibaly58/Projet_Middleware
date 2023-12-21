import json
from flask import Blueprint, request
from flask_login import login_required
from marshmallow import ValidationError

from src.models.http_exceptions import *
from src.schemas.rating import RatingSchema, RatingUpdateSchema, RatingCreateSchema
from src.schemas.errors import *
import src.services.ratings as ratings_service

# from routes import ratings
ratings = Blueprint(name="ratings", import_name=__name__)


@ratings.route('/<id>', methods=['GET'])
@login_required
def get_rating(id):
    """
    ---
    get:
      description: Getting a rating
      parameters:
        - in: path
          name: id
          schema:
            type: uuidv4
          required: true
          description: UUID of rating id
      responses:
        '200':
          description: Ok
          content:
            application/json:
              schema: Rating
            application/yaml:
              schema: Rating
        '401':
          description: Unauthorized
          content:
            application/json:
              schema: Unauthorized
            application/yaml:
              schema: Unauthorized
        '404':
          description: Not found
          content:
            application/json:
              schema: NotFound
            application/yaml:
              schema: NotFound
      tags:
          - ratings
    """
    return ratings_service.get_rating(id)


@ratings.route('/<id>', methods=['PUT'])
@login_required
def put_rating(id):
    """
    ---
    put:
      description: Updating a rating
      parameters:
        - in: path
          name: id
          schema:
            type: uuidv4
          required: true
          description: UUID of rating id
      requestBody:
        required: true
        content:
            application/json:
                schema: RatingUpdate
      responses:
        '200':
          description: Ok
          content:
            application/json:
              schema: Rating
            application/yaml:
              schema: Rating
        '401':
          description: Unauthorized
          content:
            application/json:
              schema: Unauthorized
            application/yaml:
              schema: Unauthorized
        '404':
          description: Not found
          content:
            application/json:
              schema: NotFound
            application/yaml:
              schema: NotFound
        '422':
          description: Unprocessable entity
          content:
            application/json:
              schema: UnprocessableEntity
            application/yaml:
              schema: UnprocessableEntity
      tags:
          - ratings
    """
    # parser le body
    try:
        rating_update = RatingUpdateSchema().loads(json_data=request.data.decode('utf-8'))
    except ValidationError as e:
        error = UnprocessableEntitySchema().loads(json.dumps({"message": e.messages.__str__()}))
        return error, error.get("code")

    # modification de la note (value)
    try:
        return ratings_service.modify_rating(id, rating_update)
    except UnprocessableEntity:
        error = UnprocessableEntitySchema().loads(json.dumps({"message": "One required field was empty"}))
        return error, error.get("code")
    except Forbidden:
        error = ForbiddenSchema().loads(json.dumps({"message": "Can't manage other ratings"}))
        return error, error.get("code")
    except Exception:
        error = SomethingWentWrongSchema().loads("{}")
        return error, error.get("code")


@ratings.route('/<id>', methods=['DELETE'])
@login_required
def delete_rating(id):
    """
    ---
    delete:
      description: Deleting a rating
      parameters:
        - in: path
          name: id
          schema:
            type: uuidv4
          required: true
          description: UUID of rating id
      responses:
        '204':
          description: No Content
        '401':
          description: Unauthorized
          content:
            application/json:
              schema: Unauthorized
            application/yaml:
              schema: Unauthorized
        '404':
          description: Not found
          content:
            application/json:
              schema: NotFound
            application/yaml:
              schema: NotFound
      tags:
          - ratings
    """
    try:
        ratings_service.delete_rating(id)
        return {}, 204
    except NotFound:
        error = NotFoundSchema().loads(json.dumps({"message": "Rating not found"}))
        return error, error.get("code")
    except Forbidden:
        error = ForbiddenSchema().loads(json.dumps({"message": "Can't delete other ratings"}))
        return error, error.get("code")


@ratings.route('/', methods=['GET'])
@login_required
def get_all_ratings():
    """
    ---
    get:
      description: Getting all ratings
      responses:
        '200':
          description: Ok
          content:
            application/json:
              schema: Rating
            application/yaml:
              schema: Rating
      tags:
          - ratings
    """
    return ratings_service.get_all_ratings()


@ratings.route('/', methods=['POST'])
@login_required
def create_rating():
    """
    ---
    post:
      description: Creating a new rating
      requestBody:
        required: true
        content:
            application/json:
                schema: RatingCreate
      responses:
        '201':
          description: Created
          content:
            application/json:
              schema: Rating
            application/yaml:
              schema: Rating
        '401':
          description: Unauthorized
          content:
            application/json:
              schema: Unauthorized
            application/yaml:
              schema: Unauthorized
        '422':
          description: Unprocessable entity
          content:
            application/json:
              schema: UnprocessableEntity
            application/yaml:
              schema: UnprocessableEntity
      tags:
          - ratings
    """
    try:
        # Parse the request data
        rating_create = RatingCreateSchema().loads(json_data=request.data.decode('utf-8'))
    except ValidationError as e:
        error = UnprocessableEntitySchema().loads(json.dumps({"message": e.messages.__str__()}))
        return error, error.get("code")

    # Create the rating
    try:
        return ratings_service.create_rating(rating_create)
    except Conflict:
        error = ConflictSchema().loads(json.dumps({"message": "Rating already exists"}))
        return error, error.get("code")
    except UnprocessableEntity:
        error = UnprocessableEntitySchema().loads(json.dumps({"message": "One required field was empty"}))
        return error, error.get("code")
    except Exception:
        error = SomethingWentWrongSchema().loads("{}")
        return error, error.get("code")
