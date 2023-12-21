from marshmallow import Schema, fields, validates_schema, ValidationError

class SongSchema(Schema):
    id = fields.String(description="UUID")
    title = fields.String(description="Title")
    artist = fields.String(description="Artist")

    @staticmethod
    def is_empty(obj):
        return (not obj.get("id") or obj.get("id") == "") and \
               (not obj.get("title") or obj.get("title") == "") and \
               (not obj.get("artist") or obj.get("artist") == "")

class BaseSongSchema(Schema):
    title = fields.String(description="Title")
    artist = fields.String(description="Artist")

class SongUpdateSchema(BaseSongSchema):
    @validates_schema
    def validates_schemas(self, data, **kwargs):
        if not (("title" in data and data["title"] != "") or
                ("artist" in data and data["artist"] != "")):
            raise ValidationError("at least one of ['title', 'artist'] must be specified")




class SongCreateSchema(Schema):
    title = fields.String(description="Title")
    artist = fields.String(description="Artist")
    
    # permet de définir dans quelles conditions le schéma est validé ou non
    @validates_schema
    def validates_schemas(self, data, **kwargs):
        if not any(data.get(field) for field in ["title", "artist"]):
            raise ValidationError("At least one of ['title', 'artist'] must be specified")
