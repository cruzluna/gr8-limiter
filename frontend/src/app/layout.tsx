import "./globals.css";
import type { Metadata } from "next";
import { Providers } from "./provider";
import { ClerkProvider } from "@clerk/nextjs";

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
          <Providers>{children}</Providers>
        </body>
      </html>
    </ClerkProvider>
  );
}
