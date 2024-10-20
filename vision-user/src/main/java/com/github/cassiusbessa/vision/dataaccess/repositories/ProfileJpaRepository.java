package com.github.cassiusbessa.vision.dataaccess.repositories;

import com.github.cassiusbessa.vision.dataaccess.entities.ProfileDataBaseEntity;
import org.springframework.data.jpa.repository.JpaRepository;
import org.springframework.stereotype.Component;
import org.springframework.stereotype.Repository;

import java.util.Optional;
import java.util.UUID;

@Repository
@Component
public interface ProfileJpaRepository extends JpaRepository<ProfileDataBaseEntity, UUID> {

    Optional<ProfileDataBaseEntity> findByAccountId(UUID accountId);

		Optional<ProfileDataBaseEntity> findByLink(String link);
}
