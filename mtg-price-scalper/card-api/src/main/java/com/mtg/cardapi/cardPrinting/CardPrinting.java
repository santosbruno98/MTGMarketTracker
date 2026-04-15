package com.mtg.cardapi.cardPrinting;

import com.mtg.cardapi.card.Card;
import com.mtg.cardapi.cardSet.CardSet;
import java.util.Set;

import jakarta.persistence.*;


@Entity // Related to the table cards
@Table(name = "card_printings",
        uniqueConstraints = {
            @UniqueConstraint(columnNames = {"card_id", "card_set_id", "collector_number", "is_foil"})
        }
)
public class CardPrinting {

    @Id
    @GeneratedValue(strategy = GenerationType.IDENTITY)
    private Long id;

    @ManyToOne(optional=false)
    @JoinColumn(name = "card_id")
    private Card card;

    @ManyToOne(optional=false)
    @JoinColumn(name = "set_id")
    private CardSet cardSet;

    private String collectorNumber; // 398a

    @Enumerated(EnumType.STRING)
    private Rarity rarity; // COMMON, UNCOMMON, RARE, MYTHIC

    private Boolean isFoil; // false as default

    protected CardPrinting() {
    }

    public CardPrinting(Card card, CardSet cardSet, String collectorNumber, Rarity rarity, Boolean isFoil) {
        this.card = card;   
        this.cardSet = cardSet;
        this.collectorNumber = collectorNumber;
        this.rarity = rarity;
        this.isFoil = isFoil;
    }

    // Getters and Setters
    public Long getId() { return id; }

    public Card getCard() { return card; }

    public CardSet getCardSet() { return cardSet; }

    public String getCollectorNumber() { return collectorNumber; }

    public Rarity getRarity() {
        return rarity;
    }

    public void setRarity(Rarity rarity) {
        this.rarity = rarity;
    }

    public Boolean getIsFoil() { return isFoil; }

}