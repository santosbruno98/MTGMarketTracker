package com.mtg.cardapi.cardSet;

import jakarta.persistence.*;

@Entity // Related to the table cards
@Table(name = "cardSets") // Table name
public class CardSet {

    @Id
    @GeneratedValue(strategy = GenerationType.IDENTITY)
    private Long id;

    private String codeName; // CMM
    
    protected CardSet() {
    }

    public CardSet(String codeName) {
        this.codeName = codeName;
    }

    // Getters and Setters

    public Long getId() {
        return id;
    }

    public void setId(Long id) {
        this.id = id;
    }

    public String getCodeName() {
        return codeName;
    }

    public void setCodeName(String codeName) {
        this.codeName = codeName;
    }
}