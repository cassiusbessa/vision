package com.github.cassiusbessa.vision.dataaccess.repositories;

import com.github.cassiusbessa.vision.dataaccess.entities.AccountDataBaseEntity;
import org.springframework.data.jpa.repository.JpaRepository;
import org.springframework.stereotype.Component;
import org.springframework.stereotype.Repository;

import java.util.Optional;
import java.util.UUID;

@Repository
@Component
public interface AccountJpaRepository extends JpaRepository<AccountDataBaseEntity, UUID> {
    Optional<AccountDataBaseEntity> findByEmail(String email);
}
