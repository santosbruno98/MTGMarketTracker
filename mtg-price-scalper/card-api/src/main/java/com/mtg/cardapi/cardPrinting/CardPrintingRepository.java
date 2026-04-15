package com.mtg.cardapi.cardPrinting;

import com.mtg.cardapi.card.Card;
import com.mtg.cardapi.cardSet.CardSet;
import org.springframework.data.jpa.repository.JpaRepository;

import java.util.List;

public interface CardPrintingRepository extends JpaRepository<CardPrinting, Long> {
    List<CardPrinting> findByCard(Card card);
    List<CardPrinting> findByCardSet(CardSet cardSet);
}
