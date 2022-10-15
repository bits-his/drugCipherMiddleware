-- phpMyAdmin SQL Dump
-- version 5.2.0
-- https://www.phpmyadmin.net/
--
-- Host: localhost
-- Generation Time: Oct 15, 2022 at 03:17 PM
-- Server version: 10.4.25-MariaDB
-- PHP Version: 8.1.10

SET SQL_MODE = "NO_AUTO_VALUE_ON_ZERO";
START TRANSACTION;
SET time_zone = "+00:00";


/*!40101 SET @OLD_CHARACTER_SET_CLIENT=@@CHARACTER_SET_CLIENT */;
/*!40101 SET @OLD_CHARACTER_SET_RESULTS=@@CHARACTER_SET_RESULTS */;
/*!40101 SET @OLD_COLLATION_CONNECTION=@@COLLATION_CONNECTION */;
/*!40101 SET NAMES utf8mb4 */;

--
-- Database: `drug_cipher`
--

DELIMITER $$
--
-- Procedures
--
CREATE DEFINER=`root`@`localhost` PROCEDURE `drugs` (IN `_id` INT(5), IN `query_type` VARCHAR(50), IN `_drug_name` VARCHAR(50), IN `_drug_generic_name` VARCHAR(50), IN `_dosage` INT(11), IN `_date_enrolled` TIMESTAMP, IN `_expiry_date` TIMESTAMP, IN `_nafdac` VARCHAR(100), IN `_qrcode` BLOB, IN `_manufacturers_id` INT(11))   BEGIN
IF query_type = 'insert' THEN
	INSERT INTO `drugs`(
    	drug_name,
        drug_generic_name,
        dosage,
        date_enrolled,
        expiry_date,
        nafdac,
        qr_code,
        manufacturers_id
        
    ) VALUES(
    	_drug_name,
        _drug_generic_name,
        _dosage,
        _date_enrolled,
        _expiry_date,
        _nafdac,
        _qr_code,
        _manufacturers_id
    );
    
    ELSEIF query_type ='update' THEN
    UPDATE `drugs`
    	set drug_name=_drug_name,
        drug_generic_name=_drug_generic_name,
        dosage=_dosage,
        date_enrolled=_date_enrolled,
        expiry_date=_expiry_date,
        nafdac=_nafdac,
        qr_code=_qrcode,
        Manufacturers_id=_manufacturers_id
        
    WHERE id=_id;
    
ELSEIF query_type='delete' THEN
DELETE FROM `drugs` WHERE id=_id;

END IF;

END$$

CREATE DEFINER=`root`@`localhost` PROCEDURE `Manufacturers` (IN `_id` INT(5), IN `query_type` VARCHAR(50), IN `_pharmacy_license` VARCHAR(50), IN `_pharmacist_license` VARCHAR(50), IN `_name` VARCHAR(50), IN `_address` VARCHAR(50), IN `_po_box` VARCHAR(50), IN `_email` VARCHAR(50), IN `_phone_number` VARCHAR(50))   BEGIN
IF query_type = 'insert' THEN
	INSERT INTO `Manufacturers`(
    	pharmacy_license,
        pharmacist_license,
        name,
        address,
        po_box,
        email,
        phone_number
    ) VALUES(
    	_pharmacy_license,
        _pharmacist_license,
        _name,
        _address,
        _po_box,
        _email,
        _phone_number
    );
    
    ELSEIF query_type ='update' THEN
    UPDATE `Manufacturers`
    	set pharmacy_license=_pharmacy_license,
        pharmacist_license=_pharmacist_license,
        name=_name,
        address=_address,
        po_box=_po_box,
        email=_email,
        phone_number=_phone_number
    WHERE id=_id;
    
ELSEIF query_type='delete' THEN
DELETE FROM `Manufacturers` WHERE id=_id;

END IF;

END$$

DELIMITER ;

-- --------------------------------------------------------

--
-- Table structure for table `drugs`
--

CREATE TABLE `drugs` (
  `id` int(11) NOT NULL,
  `drug_name` varchar(50) NOT NULL,
  `drug_generic_name` varchar(50) NOT NULL,
  `dosage` int(11) NOT NULL,
  `date_enrolled` timestamp NOT NULL DEFAULT current_timestamp() ON UPDATE current_timestamp(),
  `expiry_date` timestamp NOT NULL DEFAULT current_timestamp() ON UPDATE current_timestamp(),
  `nafdac` varchar(100) NOT NULL,
  `qr_code` blob NOT NULL,
  `manufacturers_id` int(11) NOT NULL
) ENGINE=InnoDB DEFAULT CHARSET=latin1;

-- --------------------------------------------------------

--
-- Table structure for table `Manufacturers`
--

CREATE TABLE `Manufacturers` (
  `manufacturers_id` int(11) NOT NULL,
  `pharmacy_license` varchar(50) NOT NULL,
  `pharmacist_license` varchar(50) NOT NULL,
  `name` varchar(50) NOT NULL,
  `address` varchar(50) NOT NULL,
  `po_box` varchar(50) NOT NULL,
  `email` varchar(50) NOT NULL,
  `phone_number` varchar(50) NOT NULL
) ENGINE=InnoDB DEFAULT CHARSET=latin1;

--
-- Indexes for dumped tables
--

--
-- Indexes for table `drugs`
--
ALTER TABLE `drugs`
  ADD PRIMARY KEY (`id`),
  ADD UNIQUE KEY `manufacturers_id` (`manufacturers_id`);

--
-- Indexes for table `Manufacturers`
--
ALTER TABLE `Manufacturers`
  ADD PRIMARY KEY (`manufacturers_id`);

--
-- Constraints for dumped tables
--

--
-- Constraints for table `drugs`
--
ALTER TABLE `drugs`
  ADD CONSTRAINT `drugs_ibfk_1` FOREIGN KEY (`manufacturers_id`) REFERENCES `Manufacturers` (`manufacturers_id`) ON DELETE CASCADE ON UPDATE CASCADE;
COMMIT;

/*!40101 SET CHARACTER_SET_CLIENT=@OLD_CHARACTER_SET_CLIENT */;
/*!40101 SET CHARACTER_SET_RESULTS=@OLD_CHARACTER_SET_RESULTS */;
/*!40101 SET COLLATION_CONNECTION=@OLD_COLLATION_CONNECTION */;
