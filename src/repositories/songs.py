from src.helpers import db
from src.models.song import Song


def get_song(song_id):
    return db.session.query(Song).filter(Song.id == song_id).first()


def get_all_songs():
    return Song.query.all()


def add_song(song):
    db.session.add(song)
    db.session.commit()


def update_song(song):
    existing_song = get_song(song.id)
    existing_song.title = song.title
    existing_song.artist = song.artist
    db.session.commit()


def delete_song(song_id):
    db.session.delete(get_song(song_id))
    db.session.commit()
