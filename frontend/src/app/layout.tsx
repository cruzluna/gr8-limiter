import "./globals.css";
import type { Metadata } from "next";
import { Providers } from "./provider";
import { ClerkProvider } from "@clerk/nextjs";
import StratusNavbar from "@/components/ui/navbar";

export const metadata: Metadata = {
  title: "stratus",
  description: "rate limiting service",
};

export default function RootLayout({
  children,
}: {
  children: React.ReactNode;
}) {
  return (
    <ClerkProvider>
      <html lang="en" className="dark">
        <body>
          <Providers>
            <section>
              <StratusNavbar />
            </section>
            {children}
          </Providers>
        </body>
      </html>
    </ClerkProvider>
  );
}
