package com.mtg.cardapi.card;

// This class will be used to resolve API calls
import org.springframework.stereotype.Service;
import java.util.List;
import java.util.Optional;

@Service
public class CardService {
    
    private final CardRepository cardRepository;
    
    public CardService(CardRepository cardRepository) {
        this.cardRepository = cardRepository;
    }
    
    public Card createCard(Card card) {
        return cardRepository.save(card);
    }

    public Optional<Card> getCard(Long id) {
        return cardRepository.findById(id);
    }

    public List<Card> searchCard(String name) {

        return cardRepository.findByNameContainingIgnoreCase(name);
    }

    public Optional<Card> updateCard(String name) {
        return cardRepository.findByNameContainingIgnoreCase(name)
                .stream()
                .findFirst()
                .map(existingCard -> {
                    existingCard.setName(name);
                    return cardRepository.save(existingCard);
                });
    }

    public boolean markToDeletion(Card cardToDelete) {
        cardToDelete.setMarkedToDeletion(true);
        cardRepository.save(cardToDelete);
        return cardToDelete.isMarkedToDeletion(); 
    }

    public void permanentDelete(Long id) throws RuntimeException {
        Card card = cardRepository.findById(id)
            .orElseThrow(() -> new CardNotFoundException(id));
        if (card.isMarkedToDeletion()) {
            cardRepository.delete(card);
        } else {
            throw new RuntimeException("This card is not marked for deletion");
        }
    }

    public Card restore(Long id) {
        Card card = cardRepository.findById(id)
        .orElseThrow(() -> new CardNotFoundException(id));
        card.setMarkedToDeletion(false);
        return cardRepository.save(card);
    }

    public Optional<Card> findByNameContainingIgnoreCase(String name) {
        return cardRepository.findByNameContainingIgnoreCase(name)
                .stream()
                .findFirst();
    }

    public List<Card> getAll() {
        return cardRepository.findAll();
    }

    public List<Card> findByMarkedToDeletion() {
        return cardRepository.findByMarkedToDeletion(true);
    }

    public Optional<Card> findById(Long id) {
        return cardRepository.findById(id);
    }
}
