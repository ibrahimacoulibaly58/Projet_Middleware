from src.helpers import db
from src.models.rating import Rating


def get_rating(id):
    return db.session.query(Rating).filter(Rating.id == id).first()


def get_ratings_for_user(user_id):
    return db.session.query(Rating).filter(Rating.user_id == user_id).all()


def add_rating(rating):
    db.session.add(rating)
    db.session.commit()


def update_rating(rating):
    existing_rating = get_rating(rating.id)
    existing_rating.rating = rating.rating
    existing_rating.comment = rating.comment
    existing_rating.rating_date = rating.rating_date
    existing_rating.song_id = rating.song_id
    existing_rating.user_id = rating.user_id
    db.session.commit()


def delete_rating(id):
    db.session.delete(get_rating(id))
    db.session.commit()
