"use client";
import { UserButton, SignInButton, SignedIn, SignedOut } from "@clerk/nextjs";
import {
  Navbar,
  NavbarBrand,
  NavbarContent,
  NavbarItem,
  Link,
  Button,
  Popover,
  PopoverTrigger,
  PopoverContent,
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
        <Popover placement="bottom" showArrow={true} color={"foreground"}>
          <PopoverTrigger>
            {/* 
// @ts-ignore */}
            <Button color={"foreground"}>
              <NavbarItem>Features</NavbarItem>
            </Button>
          </PopoverTrigger>
          <PopoverContent>
            <div className="px-1 py-2">
              <div className="text-small font-bold">
                It&apos;s in the works.
              </div>
              <div className="text-tiny">Refer to the docs for now.</div>
            </div>
          </PopoverContent>
        </Popover>

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
          <Popover placement="bottom" showArrow={true} color={"foreground"}>
            <PopoverTrigger>
              {/* 
// @ts-ignore */}
              <Button color={"foreground"}>
                <NavbarItem>Integrations</NavbarItem>
              </Button>
            </PopoverTrigger>
            <PopoverContent>
              <div className="px-1 py-2">
                <div className="text-small font-bold">
                  It&apos;s in the works.
                </div>
                <div className="text-tiny">Refer to the docs for now.</div>
              </div>
            </PopoverContent>
          </Popover>
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
