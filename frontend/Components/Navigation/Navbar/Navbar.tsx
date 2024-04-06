'use client'

import { Button, ButtonGroup } from "@chakra-ui/react";
import { Image, Link } from "@chakra-ui/next-js";

export const Navbar = () => {
    return (
        <div className={"p-4 m-8 flex align-middle"}>
            <Image src={"/assets/images/phase1.webp"} alt={"IncrediblESG Logo"} height={"60"} width={"60"} className={"rounded"}/>
            <ButtonGroup className={"p-4"}>
                <Button>
                    <Link href={"/generate"}>
                        Test
                    </Link>
                </Button>
                <Button>
                    <Link href={"/generate"}>
                        Test
                    </Link>
                </Button>
                <Button>
                    <Link href={"/generate"}>
                        Test
                    </Link>
                </Button>
                <Button>
                    <Link href={"/generate"}>
                        Test
                    </Link>
                </Button>
            </ButtonGroup>
        </div>
    );
};