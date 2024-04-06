import type { Metadata } from "next";
import { Inter } from "next/font/google";
import "./globals.css";
import { Providers } from "@/app/providers";
import { Navbar } from "@/Components/Navigation/Navbar/Navbar";
import { Form } from "@/Components/Form/Form";
import { ChakraBaseProvider, ChakraProvider } from "@chakra-ui/react";

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
              <Providers>
                <Navbar/>
                <Form />
              </Providers>
            </body>
        </html>
    );
}
