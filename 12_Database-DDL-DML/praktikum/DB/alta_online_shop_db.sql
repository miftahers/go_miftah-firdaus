-- phpMyAdmin SQL Dump
-- version 5.2.0
-- https://www.phpmyadmin.net/
--
-- Host: 127.0.0.1
-- Waktu pembuatan: 13 Sep 2022 pada 12.56
-- Versi server: 10.4.24-MariaDB
-- Versi PHP: 8.1.6

SET SQL_MODE = "NO_AUTO_VALUE_ON_ZERO";
START TRANSACTION;
SET time_zone = "+00:00";


/*!40101 SET @OLD_CHARACTER_SET_CLIENT=@@CHARACTER_SET_CLIENT */;
/*!40101 SET @OLD_CHARACTER_SET_RESULTS=@@CHARACTER_SET_RESULTS */;
/*!40101 SET @OLD_COLLATION_CONNECTION=@@COLLATION_CONNECTION */;
/*!40101 SET NAMES utf8mb4 */;

--
-- Database: `alta_online_shop_db`
--

-- --------------------------------------------------------

--
-- Struktur dari tabel `addresses`
--

CREATE TABLE `addresses` (
  `ID` int(11) NOT NULL,
  `USER_ID` int(11) NOT NULL,
  `ADDRESS` varchar(256) NOT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

-- --------------------------------------------------------

--
-- Struktur dari tabel `operators`
--

CREATE TABLE `operators` (
  `ID` int(11) NOT NULL AUTO_INCREMENT,
  `NAME` varchar(16) NOT NULL,
  `ADDRESS` varchar(256) NOT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

-- --------------------------------------------------------

--
-- Struktur dari tabel `payment_method`
--

CREATE TABLE `payment_method` (
  `ID` int(11) NOT NULL,
  `TYPE` varchar(16) NOT NULL,
  `NAME` varchar(16) NOT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

-- --------------------------------------------------------

--
-- Struktur dari tabel `products`
--

CREATE TABLE `products` (
  `ID` int(11) NOT NULL AUTO_INCREMENT,
  `OPERATOR_ID` int(11) NOT NULL,
  `NAME` varchar(64) NOT NULL,
  `PRICE` int(11) NOT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

-- --------------------------------------------------------

--
-- Struktur dari tabel `product_descriptions`
--

CREATE TABLE `product_descriptions` (
  `PRODUCT_ID` int(11) NOT NULL,
  `DESCRIPTION` text NOT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

-- --------------------------------------------------------

--
-- Struktur dari tabel `product_types`
--

CREATE TABLE `product_types` (
  `PRODUCT_ID` int(11) NOT NULL,
  `TYPE` varchar(16) NOT NULL,
  `TYPE_DESCRIPTION` varchar(32) NOT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

-- --------------------------------------------------------

--
-- Struktur dari tabel `transactions`
--

CREATE TABLE `transactions` (
  `ID` int(11) NOT NULL AUTO_INCREMENT,
  `USER_ID` int(11) NOT NULL,
  `TRANSACTION_STATUS` varchar(16) NOT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

-- --------------------------------------------------------

--
-- Struktur dari tabel `transaction_details`
--

CREATE TABLE `transaction_details` (
  `TRANSACTION_ID` int(11) NOT NULL,
  `PRODUCT_ID` int(11) NOT NULL,
  `QUANTITY` int(11) NOT NULL,
  `ORDER_DATE` date NOT NULL,
  `TOTAL_PRICE` int(11) NOT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

-- --------------------------------------------------------

--
-- Struktur dari tabel `users`
--

CREATE TABLE `users` (
  `ID` int(11) NOT NULL AUTO_INCREMENT,
  `NAME` varchar(16) NOT NULL,
  `BIRTHDAY` date DEFAULT NULL,
  `STATUS` varchar(16) DEFAULT NULL,
  `GENDER` varchar(1) NOT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

--
-- Indexes for dumped tables
--

--
-- Indeks untuk tabel `addresses`
--
ALTER TABLE `addresses`
  ADD PRIMARY KEY (`ID`),
  ADD KEY `USER_ID` (`USER_ID`);

--
-- Indeks untuk tabel `operators`
--
ALTER TABLE `operators`
  ADD PRIMARY KEY (`ID`);

--
-- Indeks untuk tabel `payment_method`
--
ALTER TABLE `payment_method`
  ADD PRIMARY KEY (`ID`);

--
-- Indeks untuk tabel `products`
--
ALTER TABLE `products`
  ADD PRIMARY KEY (`ID`),
  ADD KEY `OPERATOR_ID` (`OPERATOR_ID`);

--
-- Indeks untuk tabel `product_descriptions`
--
ALTER TABLE `product_descriptions`
  ADD KEY `PRODUCT_ID` (`PRODUCT_ID`);

--
-- Indeks untuk tabel `product_types`
--
ALTER TABLE `product_types`
  ADD KEY `PRODUCT_ID` (`PRODUCT_ID`);

--
-- Indeks untuk tabel `transactions`
--
ALTER TABLE `transactions`
  ADD PRIMARY KEY (`ID`),
  ADD KEY `USER_ID` (`USER_ID`);

--
-- Indeks untuk tabel `transaction_details`
--
ALTER TABLE `transaction_details`
  ADD KEY `TRANSACTION_ID` (`TRANSACTION_ID`),
  ADD KEY `PRODUCT_ID` (`PRODUCT_ID`);

--
-- Indeks untuk tabel `users`
--
ALTER TABLE `users`
  ADD PRIMARY KEY (`ID`);

--
-- Ketidakleluasaan untuk tabel pelimpahan (Dumped Tables)
--

--
-- Ketidakleluasaan untuk tabel `addresses`
--
ALTER TABLE `addresses`
  ADD CONSTRAINT `addresses_ibfk_1` FOREIGN KEY (`USER_ID`) REFERENCES `users` (`ID`);

--
-- Ketidakleluasaan untuk tabel `products`
--
ALTER TABLE `products`
  ADD CONSTRAINT `products_ibfk_1` FOREIGN KEY (`OPERATOR_ID`) REFERENCES `operators` (`ID`);

--
-- Ketidakleluasaan untuk tabel `product_descriptions`
--
ALTER TABLE `product_descriptions`
  ADD CONSTRAINT `product_descriptions_ibfk_1` FOREIGN KEY (`PRODUCT_ID`) REFERENCES `products` (`ID`);

--
-- Ketidakleluasaan untuk tabel `product_types`
--
ALTER TABLE `product_types`
  ADD CONSTRAINT `product_types_ibfk_1` FOREIGN KEY (`PRODUCT_ID`) REFERENCES `products` (`ID`);

--
-- Ketidakleluasaan untuk tabel `transactions`
--
ALTER TABLE `transactions`
  ADD CONSTRAINT `transactions_ibfk_1` FOREIGN KEY (`USER_ID`) REFERENCES `users` (`ID`);

--
-- Ketidakleluasaan untuk tabel `transaction_details`
--
ALTER TABLE `transaction_details`
  ADD CONSTRAINT `transaction_details_ibfk_1` FOREIGN KEY (`TRANSACTION_ID`) REFERENCES `transactions` (`ID`),
  ADD CONSTRAINT `transaction_details_ibfk_2` FOREIGN KEY (`PRODUCT_ID`) REFERENCES `products` (`ID`);
COMMIT;

/*!40101 SET CHARACTER_SET_CLIENT=@OLD_CHARACTER_SET_CLIENT */;
/*!40101 SET CHARACTER_SET_RESULTS=@OLD_CHARACTER_SET_RESULTS */;
/*!40101 SET COLLATION_CONNECTION=@OLD_COLLATION_CONNECTION */;
