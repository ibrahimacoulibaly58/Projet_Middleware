from marshmallow import Schema, fields, validates_schema, ValidationError

class RatingSchema(Schema):
    id = fields.String(description="UUID")
    comment = fields.String(description="Comment")
    rating = fields.Integer(description="Rating")
    rating_date = fields.String(description="Rating Date")
    song_id = fields.String(description="Song ID")
    user_id = fields.String(description="User ID")

    @staticmethod
    def is_empty(obj):
        return (not obj.get("id") or obj.get("id") == "") and \
               (not obj.get("comment") or obj.get("comment") == "") and \
               (obj.get("rating") is None) and \
               (not obj.get("rating_date") or obj.get("rating_date") == "") and \
               (not obj.get("song_id") or obj.get("song_id") == "") and \
               (not obj.get("user_id") or obj.get("user_id") == "")

class BaseRatingSchema(Schema):
    comment = fields.String(description="Comment")
    rating = fields.Integer(description="Rating")
    rating_date = fields.String(description="Rating Date")
    song_id = fields.String(description="Song ID")
    user_id = fields.String(description="User ID")

class RatingUpdateSchema(BaseRatingSchema):
    @validates_schema
    def validates_schemas(self, data, **kwargs):
        if not (("comment" in data and data["comment"] != "") or
                ("rating" in data and data["rating"] is not None) or
                ("rating_date" in data and data["rating_date"] != "") or
                ("song_id" in data and data["song_id"] != "") or
                ("user_id" in data and data["user_id"] != "")):
            raise ValidationError("at least one of ['comment', 'rating', 'rating_date', 'song_id', 'user_id'] must be specified")



class RatingCreateSchema(Schema):
    comment = fields.String(description="Comment", required=True)
    rating = fields.Integer(description="Rating", required=True)
    song_id = fields.String(description="UUID of the song", required=True)
    user_id = fields.String(description="UUID of the user", required=True)
    id = fields.String(description="UUID", dump_only=True)
    rating_date = fields.String(description="Rating date", dump_only=True)
    
    # Vous pouvez ajouter d'autres validations si nécessaire
    @validates_schema
    def validates_schemas(self, data, **kwargs):
        if not all(data.get(field) for field in ["comment", "rating", "song_id", "user_id"]):
            raise ValidationError("All fields must be specified")
