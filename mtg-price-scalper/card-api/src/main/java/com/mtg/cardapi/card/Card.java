package com.mtg.cardapi.card;

import java.time.LocalDateTime;

import com.mtg.cardapi.cardPrinting.CardPrinting;
import com.mtg.cardapi.cardSet.CardSet;

import jakarta.persistence.*;

import java.util.List;
import java.util.ArrayList;

// Relationships with card, set, card_printing
// Card has many card_printins and card_printings has many cards
// Set has many cards and card_printings


@Entity
@Table(name = "cards")
public class Card {

    @Id
    @GeneratedValue(strategy = GenerationType.IDENTITY)
    private Long id;
    
    private String name;
    
    private boolean markedToDeletion = Boolean.FALSE;
    
    @OneToMany(mappedBy = "card", cascade = CascadeType.ALL, orphanRemoval = true)
    private final List<CardPrinting> printings = new ArrayList<>();

    protected Card() {
    }


    public Card(String name) {
        this.name = name;
    }

    public Long getId() {
        return id;
    }

    public String getName() {
        return name;
    }

    public void setName(String name) {
        this.name = name;
    }

    public boolean isMarkedToDeletion() {
        return markedToDeletion;
    }

    public void setMarkedToDeletion(boolean markedToDeletion) {
        this.markedToDeletion = markedToDeletion;
    }
    
    public List<CardPrinting> getPrintings() {
        return printings;
    }
    
}
