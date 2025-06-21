-- phpMyAdmin SQL Dump
-- version 5.2.1
-- https://www.phpmyadmin.net/
--
-- Host: 127.0.0.1
-- Generation Time: Jun 20, 2025 at 02:42 PM
-- Server version: 10.4.32-MariaDB
-- PHP Version: 8.2.12

SET SQL_MODE = "NO_AUTO_VALUE_ON_ZERO";
START TRANSACTION;
SET time_zone = "+00:00";


/*!40101 SET @OLD_CHARACTER_SET_CLIENT=@@CHARACTER_SET_CLIENT */;
/*!40101 SET @OLD_CHARACTER_SET_RESULTS=@@CHARACTER_SET_RESULTS */;
/*!40101 SET @OLD_COLLATION_CONNECTION=@@COLLATION_CONNECTION */;
/*!40101 SET NAMES utf8mb4 */;

--
-- Database: `lppm`
--

-- --------------------------------------------------------

--
-- Table structure for table `lppm_hki_dosen`
--

CREATE TABLE `lppm_hki_dosen` (
  `id` int(5) NOT NULL,
  `judul` varchar(30) NOT NULL,
  `konten` text NOT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;

-- --------------------------------------------------------

--
-- Table structure for table `lppm_hki_mhs`
--

CREATE TABLE `lppm_hki_mhs` (
  `id` int(5) NOT NULL,
  `judul` varchar(30) NOT NULL,
  `konten` text NOT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;

-- --------------------------------------------------------

--
-- Table structure for table `lppm_jurnal_jk`
--

CREATE TABLE `lppm_jurnal_jk` (
  `id` int(5) NOT NULL,
  `judul` varchar(30) NOT NULL,
  `konten` text NOT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;

-- --------------------------------------------------------

--
-- Table structure for table `lppm_jurnal_js`
--

CREATE TABLE `lppm_jurnal_js` (
  `id` int(5) NOT NULL,
  `judul` varchar(30) NOT NULL,
  `konten` text NOT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;

-- --------------------------------------------------------

--
-- Table structure for table `lppm_jurnal_kiat`
--

CREATE TABLE `lppm_jurnal_kiat` (
  `id` int(5) NOT NULL,
  `judul` varchar(30) NOT NULL,
  `konten` text NOT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;

-- --------------------------------------------------------

--
-- Table structure for table `lppm_jurnal_tajb`
--

CREATE TABLE `lppm_jurnal_tajb` (
  `id` int(5) NOT NULL,
  `judul` varchar(30) NOT NULL,
  `konten` text NOT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;

-- --------------------------------------------------------

--
-- Table structure for table `lppm_jurnal_teknois`
--

CREATE TABLE `lppm_jurnal_teknois` (
  `id` int(5) NOT NULL,
  `judul` varchar(30) NOT NULL,
  `konten` text NOT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;

-- --------------------------------------------------------

--
-- Table structure for table `lppm_jurnal_tmjb`
--

CREATE TABLE `lppm_jurnal_tmjb` (
  `id` int(5) NOT NULL,
  `judul` varchar(30) NOT NULL,
  `konten` text NOT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;

-- --------------------------------------------------------

--
-- Table structure for table `lppm_penelitian_bame`
--

CREATE TABLE `lppm_penelitian_bame` (
  `id` int(5) NOT NULL,
  `judul` varchar(30) NOT NULL,
  `konten` text NOT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;

-- --------------------------------------------------------

--
-- Table structure for table `lppm_penelitian_hpp`
--

CREATE TABLE `lppm_penelitian_hpp` (
  `id` int(5) NOT NULL,
  `judul` varchar(30) NOT NULL,
  `konten` text NOT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;

-- --------------------------------------------------------

--
-- Table structure for table `lppm_penelitian_lp`
--

CREATE TABLE `lppm_penelitian_lp` (
  `id` int(5) NOT NULL,
  `judul` varchar(30) NOT NULL,
  `konten` text NOT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;

-- --------------------------------------------------------

--
-- Table structure for table `lppm_penelitian_p3`
--

CREATE TABLE `lppm_penelitian_p3` (
  `id` int(5) NOT NULL,
  `judul` varchar(30) NOT NULL,
  `konten` text NOT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;

-- --------------------------------------------------------

--
-- Table structure for table `lppm_penelitian_rrp`
--

CREATE TABLE `lppm_penelitian_rrp` (
  `id` int(5) NOT NULL,
  `judul` varchar(30) NOT NULL,
  `konten` text NOT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;

-- --------------------------------------------------------

--
-- Table structure for table `lppm_penelitian_sk_r`
--

CREATE TABLE `lppm_penelitian_sk_r` (
  `id` int(5) NOT NULL,
  `judul` varchar(30) NOT NULL,
  `konten` text NOT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;

-- --------------------------------------------------------

--
-- Table structure for table `lppm_penelitian_stp`
--

CREATE TABLE `lppm_penelitian_stp` (
  `id` int(5) NOT NULL,
  `judul` varchar(30) NOT NULL,
  `konten` text NOT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;

-- --------------------------------------------------------

--
-- Table structure for table `lppm_penelitian_tcr`
--

CREATE TABLE `lppm_penelitian_tcr` (
  `id` int(5) NOT NULL,
  `judul` varchar(30) NOT NULL,
  `konten` text NOT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;

-- --------------------------------------------------------

--
-- Table structure for table `lppm_pkm_bame`
--

CREATE TABLE `lppm_pkm_bame` (
  `id` int(5) NOT NULL,
  `judul` varchar(30) NOT NULL,
  `konten` text NOT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;

-- --------------------------------------------------------

--
-- Table structure for table `lppm_pkm_hpp`
--

CREATE TABLE `lppm_pkm_hpp` (
  `id` int(5) NOT NULL,
  `judul` varchar(30) NOT NULL,
  `konten` text NOT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;

-- --------------------------------------------------------

--
-- Table structure for table `lppm_pkm_lp`
--

CREATE TABLE `lppm_pkm_lp` (
  `id` int(5) NOT NULL,
  `judul` varchar(30) NOT NULL,
  `konten` text NOT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;

-- --------------------------------------------------------

--
-- Table structure for table `lppm_pkm_p3`
--

CREATE TABLE `lppm_pkm_p3` (
  `id` int(5) NOT NULL,
  `judul` varchar(30) NOT NULL,
  `konten` text NOT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;

-- --------------------------------------------------------

--
-- Table structure for table `lppm_pkm_rrp`
--

CREATE TABLE `lppm_pkm_rrp` (
  `id` int(5) NOT NULL,
  `judul` varchar(30) NOT NULL,
  `konten` text NOT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;

-- --------------------------------------------------------

--
-- Table structure for table `lppm_pkm_sk_r`
--

CREATE TABLE `lppm_pkm_sk_r` (
  `id` int(5) NOT NULL,
  `judul` varchar(30) NOT NULL,
  `konten` text NOT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;

-- --------------------------------------------------------

--
-- Table structure for table `lppm_pkm_stp`
--

CREATE TABLE `lppm_pkm_stp` (
  `id` int(5) NOT NULL,
  `judul` varchar(30) NOT NULL,
  `konten` text NOT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;

-- --------------------------------------------------------

--
-- Table structure for table `lppm_pkm_tcr`
--

CREATE TABLE `lppm_pkm_tcr` (
  `id` int(5) NOT NULL,
  `judul` varchar(30) NOT NULL,
  `konten` text NOT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;

-- --------------------------------------------------------

--
-- Table structure for table `lppm_profile_so`
--

CREATE TABLE `lppm_profile_so` (
  `id` int(5) NOT NULL,
  `judul` varchar(30) NOT NULL,
  `konten` text NOT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;

--
-- Dumping data for table `lppm_profile_so`
--

INSERT INTO `lppm_profile_so` (`id`, `judul`, `konten`) VALUES
(1, 'Struktur Organisasi LP3M', 'Kepala LP3M\r\nBertanggung jawab penuh atas perencanaan, pelaksanaan, dan evaluasi kegiatan penelitian, pengabdian kepada masyarakat, serta publikasi ilmiah.\r\n\r\nSekretaris LP3M\r\nMembantu Kepala LP3M dalam tugas administrasi, penyusunan program, serta pelaporan kegiatan LP3M.\r\n\r\nKoordinator Bidang Penelitian\r\nMengelola dan mengembangkan kegiatan penelitian dosen dan mahasiswa, termasuk pendanaan, pelatihan, dan publikasi hasil penelitian.\r\n\r\nKoordinator Bidang Pengabdian kepada Masyarakat\r\nMengkoordinasikan pelaksanaan kegiatan pengabdian kepada masyarakat berbasis hasil penelitian dan kebutuhan lokal.\r\n\r\nKoordinator Bidang Publikasi Ilmiah\r\nBertugas memfasilitasi dan meningkatkan kualitas serta kuantitas publikasi ilmiah dosen dan mahasiswa di jurnal bereputasi.\r\n\r\nKoordinator Bidang Kerja Sama dan Kemitraan\r\nMenjalin dan mengelola kerja sama dengan lembaga pemerintah, industri, dan mitra eksternal dalam bidang penelitian dan pengabdian.\r\n\r\nStaf Administrasi dan Keuangan\r\nMendukung kegiatan LP3M secara administratif dan mengelola keuangan serta pertanggungjawaban kegiatan.\r\n\r\n');

-- --------------------------------------------------------

--
-- Table structure for table `lppm_profile_visimisi`
--

CREATE TABLE `lppm_profile_visimisi` (
  `id` int(5) NOT NULL,
  `judul` varchar(30) NOT NULL,
  `konten` text NOT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;

--
-- Dumping data for table `lppm_profile_visimisi`
--

INSERT INTO `lppm_profile_visimisi` (`id`, `judul`, `konten`) VALUES
(1, 'Sistem Manajemen Website Seder', 'Visi:\r\nMembangun sistem manajemen data website yang cepat, ringan, dan efisien untuk mendukung pengelolaan informasi berbasis API secara fleksibel dan modern.\r\nMisi:\r\nMengembangkan layanan REST API yang stabil dan mudah digunakan dengan bahasa pemrograman Go.\r\n\r\nMenyediakan fitur dasar manajemen data seperti Get, Update, dan Delete dengan struktur kode yang sederhana.\r\n\r\nMendorong praktik pemrograman yang efisien dan optimal dalam pengembangan backend aplikasi.\r\n\r\nMenyediakan dasar sistem yang bisa dikembangkan lebih lanjut dengan integrasi database atau frontend antarmuka pengguna.');

--
-- Indexes for dumped tables
--

--
-- Indexes for table `lppm_hki_dosen`
--
ALTER TABLE `lppm_hki_dosen`
  ADD PRIMARY KEY (`id`);

--
-- Indexes for table `lppm_hki_mhs`
--
ALTER TABLE `lppm_hki_mhs`
  ADD PRIMARY KEY (`id`);

--
-- Indexes for table `lppm_jurnal_jk`
--
ALTER TABLE `lppm_jurnal_jk`
  ADD PRIMARY KEY (`id`);

--
-- Indexes for table `lppm_jurnal_js`
--
ALTER TABLE `lppm_jurnal_js`
  ADD PRIMARY KEY (`id`);

--
-- Indexes for table `lppm_jurnal_kiat`
--
ALTER TABLE `lppm_jurnal_kiat`
  ADD PRIMARY KEY (`id`);

--
-- Indexes for table `lppm_jurnal_tajb`
--
ALTER TABLE `lppm_jurnal_tajb`
  ADD PRIMARY KEY (`id`);

--
-- Indexes for table `lppm_jurnal_teknois`
--
ALTER TABLE `lppm_jurnal_teknois`
  ADD PRIMARY KEY (`id`);

--
-- Indexes for table `lppm_jurnal_tmjb`
--
ALTER TABLE `lppm_jurnal_tmjb`
  ADD PRIMARY KEY (`id`);

--
-- Indexes for table `lppm_penelitian_bame`
--
ALTER TABLE `lppm_penelitian_bame`
  ADD PRIMARY KEY (`id`);

--
-- Indexes for table `lppm_penelitian_hpp`
--
ALTER TABLE `lppm_penelitian_hpp`
  ADD PRIMARY KEY (`id`);

--
-- Indexes for table `lppm_penelitian_lp`
--
ALTER TABLE `lppm_penelitian_lp`
  ADD PRIMARY KEY (`id`);

--
-- Indexes for table `lppm_penelitian_p3`
--
ALTER TABLE `lppm_penelitian_p3`
  ADD PRIMARY KEY (`id`);

--
-- Indexes for table `lppm_penelitian_rrp`
--
ALTER TABLE `lppm_penelitian_rrp`
  ADD PRIMARY KEY (`id`);

--
-- Indexes for table `lppm_penelitian_sk_r`
--
ALTER TABLE `lppm_penelitian_sk_r`
  ADD PRIMARY KEY (`id`);

--
-- Indexes for table `lppm_penelitian_stp`
--
ALTER TABLE `lppm_penelitian_stp`
  ADD PRIMARY KEY (`id`);

--
-- Indexes for table `lppm_penelitian_tcr`
--
ALTER TABLE `lppm_penelitian_tcr`
  ADD PRIMARY KEY (`id`);

--
-- Indexes for table `lppm_pkm_bame`
--
ALTER TABLE `lppm_pkm_bame`
  ADD PRIMARY KEY (`id`);

--
-- Indexes for table `lppm_pkm_hpp`
--
ALTER TABLE `lppm_pkm_hpp`
  ADD PRIMARY KEY (`id`);

--
-- Indexes for table `lppm_pkm_lp`
--
ALTER TABLE `lppm_pkm_lp`
  ADD PRIMARY KEY (`id`);

--
-- Indexes for table `lppm_pkm_p3`
--
ALTER TABLE `lppm_pkm_p3`
  ADD PRIMARY KEY (`id`);

--
-- Indexes for table `lppm_pkm_rrp`
--
ALTER TABLE `lppm_pkm_rrp`
  ADD PRIMARY KEY (`id`);

--
-- Indexes for table `lppm_pkm_sk_r`
--
ALTER TABLE `lppm_pkm_sk_r`
  ADD PRIMARY KEY (`id`);

--
-- Indexes for table `lppm_pkm_stp`
--
ALTER TABLE `lppm_pkm_stp`
  ADD PRIMARY KEY (`id`);

--
-- Indexes for table `lppm_pkm_tcr`
--
ALTER TABLE `lppm_pkm_tcr`
  ADD PRIMARY KEY (`id`);

--
-- Indexes for table `lppm_profile_so`
--
ALTER TABLE `lppm_profile_so`
  ADD PRIMARY KEY (`id`);

--
-- Indexes for table `lppm_profile_visimisi`
--
ALTER TABLE `lppm_profile_visimisi`
  ADD PRIMARY KEY (`id`);

--
-- AUTO_INCREMENT for dumped tables
--

--
-- AUTO_INCREMENT for table `lppm_hki_dosen`
--
ALTER TABLE `lppm_hki_dosen`
  MODIFY `id` int(5) NOT NULL AUTO_INCREMENT;

--
-- AUTO_INCREMENT for table `lppm_hki_mhs`
--
ALTER TABLE `lppm_hki_mhs`
  MODIFY `id` int(5) NOT NULL AUTO_INCREMENT;

--
-- AUTO_INCREMENT for table `lppm_jurnal_jk`
--
ALTER TABLE `lppm_jurnal_jk`
  MODIFY `id` int(5) NOT NULL AUTO_INCREMENT;

--
-- AUTO_INCREMENT for table `lppm_jurnal_js`
--
ALTER TABLE `lppm_jurnal_js`
  MODIFY `id` int(5) NOT NULL AUTO_INCREMENT;

--
-- AUTO_INCREMENT for table `lppm_jurnal_kiat`
--
ALTER TABLE `lppm_jurnal_kiat`
  MODIFY `id` int(5) NOT NULL AUTO_INCREMENT;

--
-- AUTO_INCREMENT for table `lppm_jurnal_tajb`
--
ALTER TABLE `lppm_jurnal_tajb`
  MODIFY `id` int(5) NOT NULL AUTO_INCREMENT;

--
-- AUTO_INCREMENT for table `lppm_jurnal_teknois`
--
ALTER TABLE `lppm_jurnal_teknois`
  MODIFY `id` int(5) NOT NULL AUTO_INCREMENT;

--
-- AUTO_INCREMENT for table `lppm_jurnal_tmjb`
--
ALTER TABLE `lppm_jurnal_tmjb`
  MODIFY `id` int(5) NOT NULL AUTO_INCREMENT;

--
-- AUTO_INCREMENT for table `lppm_penelitian_bame`
--
ALTER TABLE `lppm_penelitian_bame`
  MODIFY `id` int(5) NOT NULL AUTO_INCREMENT;

--
-- AUTO_INCREMENT for table `lppm_penelitian_hpp`
--
ALTER TABLE `lppm_penelitian_hpp`
  MODIFY `id` int(5) NOT NULL AUTO_INCREMENT;

--
-- AUTO_INCREMENT for table `lppm_penelitian_lp`
--
ALTER TABLE `lppm_penelitian_lp`
  MODIFY `id` int(5) NOT NULL AUTO_INCREMENT;

--
-- AUTO_INCREMENT for table `lppm_penelitian_p3`
--
ALTER TABLE `lppm_penelitian_p3`
  MODIFY `id` int(5) NOT NULL AUTO_INCREMENT;

--
-- AUTO_INCREMENT for table `lppm_penelitian_rrp`
--
ALTER TABLE `lppm_penelitian_rrp`
  MODIFY `id` int(5) NOT NULL AUTO_INCREMENT;

--
-- AUTO_INCREMENT for table `lppm_penelitian_sk_r`
--
ALTER TABLE `lppm_penelitian_sk_r`
  MODIFY `id` int(5) NOT NULL AUTO_INCREMENT;

--
-- AUTO_INCREMENT for table `lppm_penelitian_stp`
--
ALTER TABLE `lppm_penelitian_stp`
  MODIFY `id` int(5) NOT NULL AUTO_INCREMENT;

--
-- AUTO_INCREMENT for table `lppm_penelitian_tcr`
--
ALTER TABLE `lppm_penelitian_tcr`
  MODIFY `id` int(5) NOT NULL AUTO_INCREMENT;

--
-- AUTO_INCREMENT for table `lppm_pkm_bame`
--
ALTER TABLE `lppm_pkm_bame`
  MODIFY `id` int(5) NOT NULL AUTO_INCREMENT;

--
-- AUTO_INCREMENT for table `lppm_pkm_hpp`
--
ALTER TABLE `lppm_pkm_hpp`
  MODIFY `id` int(5) NOT NULL AUTO_INCREMENT;

--
-- AUTO_INCREMENT for table `lppm_pkm_lp`
--
ALTER TABLE `lppm_pkm_lp`
  MODIFY `id` int(5) NOT NULL AUTO_INCREMENT;

--
-- AUTO_INCREMENT for table `lppm_pkm_p3`
--
ALTER TABLE `lppm_pkm_p3`
  MODIFY `id` int(5) NOT NULL AUTO_INCREMENT;

--
-- AUTO_INCREMENT for table `lppm_pkm_rrp`
--
ALTER TABLE `lppm_pkm_rrp`
  MODIFY `id` int(5) NOT NULL AUTO_INCREMENT;

--
-- AUTO_INCREMENT for table `lppm_pkm_sk_r`
--
ALTER TABLE `lppm_pkm_sk_r`
  MODIFY `id` int(5) NOT NULL AUTO_INCREMENT;

--
-- AUTO_INCREMENT for table `lppm_pkm_stp`
--
ALTER TABLE `lppm_pkm_stp`
  MODIFY `id` int(5) NOT NULL AUTO_INCREMENT;

--
-- AUTO_INCREMENT for table `lppm_pkm_tcr`
--
ALTER TABLE `lppm_pkm_tcr`
  MODIFY `id` int(5) NOT NULL AUTO_INCREMENT;

--
-- AUTO_INCREMENT for table `lppm_profile_so`
--
ALTER TABLE `lppm_profile_so`
  MODIFY `id` int(5) NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=2;

--
-- AUTO_INCREMENT for table `lppm_profile_visimisi`
--
ALTER TABLE `lppm_profile_visimisi`
  MODIFY `id` int(5) NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=2;
COMMIT;

/*!40101 SET CHARACTER_SET_CLIENT=@OLD_CHARACTER_SET_CLIENT */;
/*!40101 SET CHARACTER_SET_RESULTS=@OLD_CHARACTER_SET_RESULTS */;
/*!40101 SET COLLATION_CONNECTION=@OLD_COLLATION_CONNECTION */;
