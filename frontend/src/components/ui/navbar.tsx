"use client";
import { UserButton, SignInButton, SignedIn, SignedOut } from "@clerk/nextjs";
import {
  Navbar,
  NavbarBrand,
  NavbarContent,
  NavbarItem,
  Link,
  Button,
} from "@nextui-org/react";
import { useRouter } from "next/navigation";

export default function StratusNavbar() {
  const router = useRouter();
  return (
    <Navbar>
      <NavbarBrand
        onClick={() => {
          router.push("/");
        }}
      >
        <p className="font-light text-xl text-inherit">🌩️ stratus</p>
      </NavbarBrand>
      <NavbarContent className="hidden sm:flex gap-4" justify="center">
        <NavbarItem>
          <Link color="foreground" href="#">
            Features
          </Link>
        </NavbarItem>
        <NavbarItem isActive>
          <Link
            href="https://stratus-docs-production.up.railway.app/"
            target="_blank"
            aria-current="page"
          >
            Docs
          </Link>
        </NavbarItem>
        <NavbarItem>
          <Link color="foreground" href="#">
            Integrations
          </Link>
        </NavbarItem>
      </NavbarContent>
      <NavbarContent justify="end">
        <SignedOut>
          <NavbarItem className="lg:flex">
            <Link href="/sign-in">Login</Link>
          </NavbarItem>
          <NavbarItem>
            <Button as={Link} color="primary" href="/sign-up" variant="flat">
              Sign Up
            </Button>
          </NavbarItem>
        </SignedOut>

        <SignedIn>
          <NavbarItem>
            <Button as={Link} color="primary" href="/dashboard" variant="flat">
              Dashboard
            </Button>
          </NavbarItem>
        </SignedIn>
      </NavbarContent>
      <SignedIn>
        {/* Mount the UserButton component */}
        <UserButton />
      </SignedIn>
    </Navbar>
  );
}
