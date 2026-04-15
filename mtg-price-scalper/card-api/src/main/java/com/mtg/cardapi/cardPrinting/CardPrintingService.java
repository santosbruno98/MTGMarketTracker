package com.mtg.cardapi.cardPrinting;

import com.mtg.cardapi.card.Card;
import com.mtg.cardapi.cardPrinting.CardPrintingRepository;
import com.mtg.cardapi.cardSet.CardSet;
import org.springframework.stereotype.Service;

import java.util.List;
import java.util.Optional;

@Service
public class CardPrintingService {
    private final CardPrintingRepository cardPrintingRepository;

    public CardPrintingService(CardPrintingRepository cardPrintingRepository) {
        this.cardPrintingRepository = cardPrintingRepository;
    }

    public CardPrinting createPrinting(CardPrinting printing) {
        return cardPrintingRepository.save(printing);
    }

    public Optional<CardPrinting> findById(Long id){
        return cardPrintingRepository.findById(id);
    }

    public List<CardPrinting> findByCard(Card card){
        return cardPrintingRepository.findByCard(card);
    }

    public List<CardPrinting> findByCardSet(CardSet cardSet) {
        return cardPrintingRepository.findByCardSet(cardSet);
    }

    public void permanentDelete(Long id) {
        cardPrintingRepository.deleteById(id);
    }
}
