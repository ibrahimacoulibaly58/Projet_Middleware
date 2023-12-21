from src.helpers import db

class Rating(db.Model):
    __tablename__ = 'ratings'

    id = db.Column(db.String(255), primary_key=True)
    user_id = db.Column(db.String(255), nullable=False)
    song_id = db.Column(db.String(255), nullable=False)
    rating = db.Column(db.Integer, nullable=False)
    comment = db.Column(db.String, nullable=True)
    rating_date = db.Column(db.String, nullable=True)

    def __init__(self, rating_id, user_id, song_id, rating, comment=None, rating_date=None):
        self.id = rating_id
        self.user_id = user_id
        self.song_id = song_id
        self.rating = rating
        self.comment = comment
        self.rating_date = rating_date

    def is_empty(self):
        return (not self.id or self.id == "") and \
               (not self.user_id or self.user_id == "") and \
               (not self.song_id or self.song_id == "") and \
               (not self.rating or self.rating == "")
