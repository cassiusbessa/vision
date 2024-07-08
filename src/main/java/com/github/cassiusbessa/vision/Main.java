package com.github.cassiusbessa.vision;

import org.springframework.boot.SpringApplication;
import org.springframework.boot.autoconfigure.EnableAutoConfiguration;
import org.springframework.boot.autoconfigure.SpringBootApplication;
import org.springframework.boot.autoconfigure.domain.EntityScan;
import org.springframework.boot.autoconfigure.security.servlet.SecurityAutoConfiguration;
import org.springframework.context.annotation.ComponentScan;


//@ComponentScan(basePackages = {"com.github.cassiusbessa.vision.*", "com.github.cassiusbessa.vision"})
//@EntityScan(basePackages = {"com.github.cassiusbessa.vision.*", "com.github.cassiusbessa.vision"})
//@EnableAutoConfiguration
@SpringBootApplication( exclude = { SecurityAutoConfiguration.class })
public class Main {

    public static void main(String[] args) {
        SpringApplication.run(Main.class, args);
    }
}