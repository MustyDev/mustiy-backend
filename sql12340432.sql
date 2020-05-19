-- phpMyAdmin SQL Dump
-- version 4.7.1
-- https://www.phpmyadmin.net/
--
-- Host: sql12.freemysqlhosting.net
-- Generation Time: May 19, 2020 at 01:21 PM
-- Server version: 5.5.62-0ubuntu0.14.04.1
-- PHP Version: 7.0.33-0ubuntu0.16.04.3

SET SQL_MODE = "NO_AUTO_VALUE_ON_ZERO";
SET AUTOCOMMIT = 0;
START TRANSACTION;
SET time_zone = "+00:00";


/*!40101 SET @OLD_CHARACTER_SET_CLIENT=@@CHARACTER_SET_CLIENT */;
/*!40101 SET @OLD_CHARACTER_SET_RESULTS=@@CHARACTER_SET_RESULTS */;
/*!40101 SET @OLD_COLLATION_CONNECTION=@@COLLATION_CONNECTION */;
/*!40101 SET NAMES utf8mb4 */;

--
-- Database: `sql12340432`
--

-- --------------------------------------------------------

--
-- Table structure for table `donasi`
--

CREATE TABLE `donasi` (
  `id` int(11) NOT NULL,
  `judul` varchar(255) NOT NULL,
  `kategori` int(9) NOT NULL,
  `nama` varchar(255) NOT NULL,
  `organisasi` varchar(255) NOT NULL,
  `email` varchar(255) NOT NULL,
  `nominal` int(222) NOT NULL,
  `deskripsi` varchar(1000) NOT NULL,
  `waktu_start` varchar(255) NOT NULL,
  `waktu_end` varchar(255) NOT NULL,
  `url` varchar(255) NOT NULL,
  `status` varchar(200) NOT NULL,
  `jumlah` int(222) NOT NULL
) ENGINE=InnoDB DEFAULT CHARSET=latin1;

--
-- Dumping data for table `donasi`
--

INSERT INTO `donasi` (`id`, `judul`, `kategori`, `nama`, `organisasi`, `email`, `nominal`, `deskripsi`, `waktu_start`, `waktu_end`, `url`, `status`, `jumlah`) VALUES
(18, 'Konser JAZZ Indonesia', 1, 'Bayu', 'JAZZ INDONESIA', 'bayu@gmail.com', 10000000, 'Kami akan mengadakan Konser JAZZ di daerah Yogyakarta dalam rangka membantu para pasien dan juga tenaga medis di Yogyakarta yang sudah berjuang selama ini demi melawan COVID-19, tetapi dikarenakan kurangnya dana untuk mengadakan Konser Amal tersebut kami meminta bantuan dari semua penggemar musik JAZZ demi terlaksananya Konser Amal tersebut', '2020-05-19', '2020-05-31', 'https://magazine.yoexplore.co.id/wp-content/uploads/2018/07/festival-jazz-indonesia-2018-YoExplore-nanaimojazzfest.jpg', 'approve', 0),
(19, 'Rock Ring Indonesia', 1, 'Enggar', 'Persatuan Musisi Rock Indonesia', 'enggar@gmail.com', 20000000, 'Kami akan mengadakan Rock Ring di daerah Surabaya dalam rangka membantu para pasien dan juga tenaga medis di Surabaya yang sudah berjuang selama ini demi melawan COVID-19, tetapi dikarenakan kurangnya dana untuk mengadakan Konser Amal tersebut kami meminta bantuan dari semua penggemar musik ROCK demi terlaksananya Konser Amal tersebut', '2020-05-20', '2020-06-30', 'https://music.mxdwn.com/wp-content/uploads/2014/10/Rock-am-Ring-logo.jpg', 'approve', 0),
(25, 'Galang Dana Konser Sheila On 7', 2, 'dsa', 'Ismaya', 'musty@mustyc.com', 22222222, 'asdasd', '2020-05-26', '2020-05-30', 'https://blue.kumparan.com/image/upload/fl_progressive,fl_lossy,c_fill,q_auto:best,w_640/v1586691973/egwky7remuvtiitdtnhb.png', 'approve', 250000),
(26, 'Galang Dana Konser Sheila On 7', 2, 'dsa', 'Ismaya', 'musty@mustyc.com', 2147483647, 'asdsad', '2020-05-11', '2020-05-23', 'https://blue.kumparan.com/image/upload/fl_progressive,fl_lossy,c_fill,q_auto:best,w_640/v1586691973/egwky7remuvtiitdtnhb.png', 'approve', 0),
(27, 'Galang Dana Konser Sheila On 7', 3, 'TariUgi', 'Ismaya', 'gentakamsa@gmail.com', 2147483647, 'asdsad', '2020-05-20', '2020-05-30', 'https://blue.kumparan.com/image/upload/fl_progressive,fl_lossy,c_fill,q_auto:best,w_640/v1586691973/egwky7remuvtiitdtnhb.png', 'approve', 0),
(28, 'Galang Dana Konser Sheila On 7', 2, 'TariUgi', 'Ismaya', 'musty@mustyc.com', 22222222, 'sadad', '2020-05-20', '2020-05-30', 'https://blue.kumparan.com/image/upload/fl_progressive,fl_lossy,c_fill,q_auto:best,w_640/v1586691973/egwky7remuvtiitdtnhb.png', 'approve', 0),
(29, 'Galang Dana Konser Sheila On 7', 1, 'Enggarr', 'Ismaya', 'enggarseptrinas@rocketmail.comss', 33333333, 'asdasd', '2020-05-24', '2020-05-30', 'https://blue.kumparan.com/image/upload/fl_progressive,fl_lossy,c_fill,q_auto:best,w_640/v1586691973/egwky7remuvtiitdtnhb.png', 'waiting', 0);

-- --------------------------------------------------------

--
-- Table structure for table `users`
--

CREATE TABLE `users` (
  `id` int(32) NOT NULL,
  `Username` varchar(32) NOT NULL,
  `Email` varchar(32) NOT NULL,
  `Phone` varchar(32) NOT NULL,
  `Password` varchar(255) NOT NULL,
  `Role` int(8) NOT NULL,
  `Token` varchar(100) NOT NULL
) ENGINE=InnoDB DEFAULT CHARSET=latin1;

--
-- Dumping data for table `users`
--

INSERT INTO `users` (`id`, `Username`, `Email`, `Phone`, `Password`, `Role`, `Token`) VALUES
(61, 'Admin', 'admin@musty.com', '08123456789', '$2a$10$Sis8RjG9XRfCksFJD2z6Qu1AepYxHk6JatyMc8uTVo9iFRTktNrR6', 0, 'secret'),
(62, 'EnggarSe', 'enggarseptrinas@musty.com', '08126700585', '$2a$10$6kWxiqW4kezWzSoS1AQ9eeQbuy5C9kbL3gUNLyq2FH4ZlfmQ0zyqm', 1, 'secret'),
(63, 'Hadyd', 'hadyd@gmail.com', '08123456789', '$2a$10$nqBg.GVQBDODQosAjgwXzuDv4JT874HtbXl5Ij6M9zXpFVNRYhV9e', 1, 'secret'),
(64, 'Musisi', 'musisi@musty.com', '08121963555', '$2a$10$DL06smBUmOA45iumgXmDVuGFy3SLPX92U0unotM4TxVa9Y2ALh7Da', 1, 'secret');

--
-- Indexes for dumped tables
--

--
-- Indexes for table `donasi`
--
ALTER TABLE `donasi`
  ADD PRIMARY KEY (`id`);

--
-- Indexes for table `users`
--
ALTER TABLE `users`
  ADD PRIMARY KEY (`id`);

--
-- AUTO_INCREMENT for dumped tables
--

--
-- AUTO_INCREMENT for table `donasi`
--
ALTER TABLE `donasi`
  MODIFY `id` int(11) NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=30;
--
-- AUTO_INCREMENT for table `users`
--
ALTER TABLE `users`
  MODIFY `id` int(32) NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=65;COMMIT;

/*!40101 SET CHARACTER_SET_CLIENT=@OLD_CHARACTER_SET_CLIENT */;
/*!40101 SET CHARACTER_SET_RESULTS=@OLD_CHARACTER_SET_RESULTS */;
/*!40101 SET COLLATION_CONNECTION=@OLD_COLLATION_CONNECTION */;
