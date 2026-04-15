package com.mtg.cardapi.card;
//https://www.baeldung.com/spring-requestmapping
import org.springframework.web.bind.annotation.DeleteMapping;
import org.springframework.web.bind.annotation.GetMapping;
import org.springframework.web.bind.annotation.PathVariable;
import org.springframework.web.bind.annotation.PostMapping;
import org.springframework.web.bind.annotation.PutMapping;
import org.springframework.web.bind.annotation.RequestBody;
import org.springframework.web.bind.annotation.RestController;
import org.springframework.web.bind.annotation.RequestParam;
import org.springframework.web.bind.annotation.RequestMapping;
import java.util.List;

@RestController
@RequestMapping("/api/cards")
public class CardController {
    //inject the service
    private final CardService cardService; // uses methods to update the bd

    CardController(CardService cardService) {
        this.cardService = cardService;
    }

    @GetMapping("/{id}")
    public Card getCardById(@PathVariable Long id) {
        return cardService.findById(id)
                .orElseThrow(() -> new CardNotFoundException(id));
    }

    @PutMapping("/{name}")
    public Card updateCard(@PathVariable String name) {
        // find the card them update it

        return cardService.updateCard(name)
                .orElseThrow(() -> new CardNotFoundException(name));
    }

    @DeleteMapping("/{id}")
    public boolean markForDeletion(@PathVariable Long id) {
        Card cardMarkedToDelete = cardService.findById(id)
        .orElseThrow(() -> new RuntimeException("Card not found"));

        return cardService.markToDeletion(cardMarkedToDelete);
    }

    @DeleteMapping("/{id}/permanent")
    public void permanentDelete(@PathVariable Long id) {
        cardService.permanentDelete(id);
    }

    @PutMapping("/{id}/restore")
    public Card restore(@PathVariable Long id) {
        return cardService.restore(id);
    }

    @GetMapping("/search")
    public Card searchCard(@RequestParam String name) {
        return cardService.findByNameContainingIgnoreCase(name)
                .orElseThrow(() -> new CardNotFoundException(name));
    }

    @PostMapping
    public Card createCard(@RequestBody Card card) {
        return cardService.createCard(card);
    }

    @GetMapping("/marked-for-deletion")
    public List<Card> getMarkedForDeletion() {
        return cardService.findByMarkedToDeletion();
    }
}
