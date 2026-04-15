package com.mtg.cardapi.card;

// This class will be used to access the database
import org.springframework.data.jpa.repository.JpaRepository;
import java.util.List;

public interface CardRepository extends JpaRepository<Card, Long> {
    List<Card> findByNameContainingIgnoreCase(String name);
    List<Card> findByMarkedToDeletion(boolean markedToDeletion);

}
