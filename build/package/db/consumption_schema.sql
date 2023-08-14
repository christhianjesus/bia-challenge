DROP TABLE IF EXISTS consumptions;
CREATE TABLE consumptions (
  id TEXT,
  meter_id INT,
  active_energy DOUBLE PRECISION,
  reactive_energy DOUBLE PRECISION,
  capacitive_reactive DOUBLE PRECISION,
  solar DOUBLE PRECISION,
  "date" TIMESTAMP
);

CREATE INDEX ON consumptions (meter_id);
CREATE INDEX ON consumptions ("date");
