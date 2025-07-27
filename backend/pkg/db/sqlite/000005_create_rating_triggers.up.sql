CREATE TRIGGER update_rating_after_insert
AFTER INSERT ON ratings
BEGIN
  UPDATE ads
  SET
    average_rating = (
      SELECT ROUND(AVG(rating), 2)
      FROM ratings
      WHERE ad_id = NEW.ad_id
    ),
    ratings_count = (
      SELECT COUNT(*)
      FROM ratings
      WHERE ad_id = NEW.ad_id
    )
  WHERE id = NEW.ad_id;
END;

CREATE TRIGGER update_rating_after_update
AFTER UPDATE ON ratings
BEGIN
  UPDATE ads
  SET
    average_rating = (
      SELECT ROUND(AVG(rating), 2)
      FROM ratings
      WHERE ad_id = NEW.ad_id
    ),
    ratings_count = (
      SELECT COUNT(*)
      FROM ratings
      WHERE ad_id = NEW.ad_id
    )
  WHERE id = NEW.ad_id;
END;

CREATE TRIGGER update_rating_after_delete
AFTER DELETE ON ratings
BEGIN
  UPDATE ads
  SET
    average_rating = (
      SELECT IFNULL(ROUND(AVG(rating), 2), 0)
      FROM ratings
      WHERE ad_id = OLD.ad_id
    ),
    ratings_count = (
      SELECT COUNT(*)
      FROM ratings
      WHERE ad_id = OLD.ad_id
    )
  WHERE id = OLD.ad_id;
END;
