package com.mtg.cardapi.card;

import org.springframework.http.HttpStatus;
import org.springframework.web.bind.annotation.ResponseStatus;

@ResponseStatus(HttpStatus.NOT_FOUND)
class CardNotFoundException extends RuntimeException{
    CardNotFoundException(Long id) {

        super("Card with id " + id + " not found");
    }
    CardNotFoundException(String name) {
        super("Card with name " + name + " not found");
    }
    CardNotFoundException(String name, Throwable cause) {
        super("Card with name '"+name+"' not found", cause);
    }
}
