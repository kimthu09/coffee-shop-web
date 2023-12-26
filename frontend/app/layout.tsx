import type { Metadata } from "next";
import { Inter } from "next/font/google";
import "./globals.css";
import Header from "@/components/header";
import HeaderMobile from "@/components/header-mobile";
import { SidebarProvider, SidebarNav } from "@/components/sidebar";
import { Toaster } from "@/components/ui/toaster";
import { AuthProvider } from "@/hooks/auth-context";

const inter = Inter({ subsets: ["latin"] });

export const metadata: Metadata = {
  title: "CoffeeShop",
  description: "Coffeeshop management app",
  icons: {
    icon: ["/favicon.ico?v=4"],
    apple: ["/apple-touch-icon.png?v=4"],
    shortcut: ["apple-touch-icon.png"],
  },
  manifest: "/site.webmanifest",
};
export default function RootLayout({
  children,
}: {
  children: React.ReactNode;
}) {
  return (
    <html lang="en" className="h-full">
      <body className={`${inter.className} flex overflow-y-hidden h-full`}>
        <AuthProvider>
          <SidebarProvider>
            <SidebarNav />
          </SidebarProvider>

          <main className="flex flex-1">
            <div className="flex flex-1 flex-col overflow-y-hidden">
              <Header />
              <HeaderMobile />
              <div className="md:p-8 p-4 overflow-auto min-w-0 ">
                {children}
              </div>
              <Toaster />
            </div>
          </main>
        </AuthProvider>
      </body>
    </html>
  );
}
