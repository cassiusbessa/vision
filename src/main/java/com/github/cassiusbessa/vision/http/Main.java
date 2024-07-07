package com.github.cassiusbessa.vision.http;

import org.springframework.boot.SpringApplication;
import org.springframework.boot.autoconfigure.SpringBootApplication;
import org.springframework.context.annotation.ComponentScan;

@SpringBootApplication
@ComponentScan(basePackages = {"com.github.cassiusbessa.vision.*", "com.github.cassiusbessa.vision"})
public class Main {

    public static void main(String[] args) {
        SpringApplication.run(Main.class, args);
    }
}