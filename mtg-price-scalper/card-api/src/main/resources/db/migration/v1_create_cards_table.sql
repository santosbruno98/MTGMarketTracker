CREATE TABLE cards IF NOT EXISTS (
    id UUID PRIMARY KEY,
    name VARCHAR(255) NOT NULL,
    mana_cost TEXT NOT NULL,
    is_comander BOOLEAN NOT NULL,
    flavour_text TEXT NOT NULL,
)

CREATE TABLE types IF NOT EXISTS ( -- Permanents -> Lands, Creatures, Artifacts, Enchament, Planeswalker, Battle, Saga 
    id UUID PRIMARY KEY,
    name TEXT NOT NULL,
    description TEXT NOT NULL,
)

CREATE TABLE supertypes IF NOT EXISTS (
    id UUID PRIMARY KEY,
    name TEXT NOT NULL,
)

CREATE TABLE subtypes IF NOT EXISTS (
    id UUID PRIMARY KEY,
    name TEXT NOT NULL,
)

CREATE TABLE card_types IF NOT EXISTS (
    id UUID PRIMARY KEY,
    type_id UUID NOT NULL,
    supertype_id UUID NOT NULL,
    subtype_id UUID NOT NULL,
    PRIMARY KEY (type_id, supertype_id, subtype_id),
)

CREATE TABLE card_supertypes IF NOT EXISTS (
    id UUID PRIMARY KEY,
    supertype_id UUID NOT NULL,
    PRIMARY KEY (supertype_id),
)

CREATE TABLE card_subtypes IF NOT EXISTS (
    id UUID PRIMARY KEY,
    subtype_id UUID NOT NULL,
    PRIMARY KEY (subtype_id),
)

CREATE TABLE sets IF NOT EXISTS (
    id UUID PRIMARY KEY,
    name TEXT NOT NULL,
    code TEXT NOT NULL,
    released BOOLEAN NOT NULL,
    block TEXT NOT NULL,
    type_id UUID NOT NULL,
    PRIMARY KEY (type_id),
)

CREATE TABLE card_sets IF NOT EXISTS (
    id UUID PRIMARY KEY,
    set_id UUID NOT NULL,
    PRIMARY KEY (set_id),
)

