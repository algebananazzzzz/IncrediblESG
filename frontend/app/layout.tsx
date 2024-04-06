import type { Metadata } from "next";
import { Inter } from "next/font/google";
import "./globals.css";
import { Providers } from "@/app/providers";
import { Navbar } from "@/Components/Navigation/Navbar/Navbar";

const inter = Inter({ subsets: ["latin"] });

export const metadata: Metadata = {
  title: "IncrediblESG",
  description: "ESG Metric Dashboard",
};

export default function RootLayout({
  children,
}: Readonly<{
  children: React.ReactNode;
}>) {
    return(
        <html lang="en">
            <body className={inter.className}>
                <Navbar/>
                <Providers>
                    {children}
                </Providers>
            </body>
        </html>
    );
}
