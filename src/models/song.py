from src.helpers import db

class Song(db.Model):
    __tablename__ = 'songs'

    id = db.Column(db.String(255), primary_key=True)
    title = db.Column(db.String(255), nullable=False)
    artist = db.Column(db.String(255), nullable=False)

    def __init__(self, song_id, title, artist):
        self.id = song_id
        self.title = title
        self.artist = artist

    def is_empty(self):
        return (not self.id or self.id == "") and \
               (not self.title or self.title == "") and \
               (not self.artist or self.artist == "")
