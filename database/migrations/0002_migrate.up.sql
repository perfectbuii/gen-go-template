-- Just make sure you already enabled PGroonga Extension
CREATE EXTENSION IF NOT EXISTS pgroonga;

CREATE INDEX IF NOT EXISTS team_pgroonga_type_index ON teams USING pgroonga (type);

CREATE INDEX IF NOT EXISTS team_pgroonga_name_index ON teams USING pgroonga (name);

CREATE INDEX IF NOT EXISTS hub_pgroonga_name_index ON hubs USING pgroonga (name);