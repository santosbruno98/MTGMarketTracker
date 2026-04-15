package com.mtg.cardapi.cardPrinting;

import org.springframework.web.bind.annotation.GetMapping;
import org.springframework.web.bind.annotation.RequestMapping;
import org.springframework.web.bind.annotation.RestController;

@RestController
@RequestMapping("/api/card-sets")
public class CardPrintingController {

    private final CardPrintingService cardPrintingService;

    CardPrintingController(CardPrintingService cardPrintingService) {
        this.cardPrintingService = cardPrintingService;
    }

    @GetMapping("/{codeSet}")
}
