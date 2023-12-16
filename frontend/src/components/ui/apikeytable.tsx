"use client";
import {
  Button,
  Chip,
  Kbd,
  Listbox,
  ListboxItem,
  ListboxSection,
  Modal,
  ModalBody,
  ModalContent,
  ModalFooter,
  ModalHeader,
  Spacer,
  Table,
  TableBody,
  TableCell,
  TableColumn,
  TableHeader,
  TableRow,
  Tooltip,
  useDisclosure,
} from "@nextui-org/react";
import { AddNoteIcon, ClipboardIcon, DeleteIcon } from "./icons";
import { ApiKeyPayload } from "@/app/(user)/dashboard/page";
import { v4 as uuidv4 } from "uuid";
import toast, { Toaster } from "react-hot-toast";
import { useState } from "react";

type ApiKeyTableProps = {
  userId: string;
  apiKeyData: ApiKeyPayload[];
};

const columns = [
  { name: "API KEY", uid: "apikey" },
  { name: "STATUS", uid: "status" },
  { name: "CREATED ON", uid: "date" },
  { name: "ACTIONS", uid: "actions" },
];

export default function ApiKeyTable({ userId, apiKeyData }: ApiKeyTableProps) {
  const iconClasses =
    "text-xl text-default-500 pointer-events-none flex-shrink-0";

  const GenerateAndInsertApiKey = async () => {
    await fetch("/api/apikey", {
      method: "POST",
      body: JSON.stringify({
        api_key: uuidv4(),

        user_id: userId,
      }),
      headers: {
        "Content-Type": "application/json",
      },
    })
      .then((res) => {
        if (res.status == 429) {
          // rate limited or too many api keys already
          res.text().then((message) => {
            toast.error(message);
          });
        } else if (res.status == 200) {
          toast.success("Successfully generated API key.");
        } else {
          res.text().then((message) => {
            toast.error(message);
          });
        }
      })
      .catch((error) => {
        toast.error("Error: ", error);
      });
  };
  const DeleteApiKey = async () => {
    if (!apiKeyToDelete || typeof apiKeyToDelete !== "string") {
      toast.error("Unable to delete api key");
    }
    await fetch("/api/apikey", {
      method: "DELETE",
      body: JSON.stringify({
        api_key: apiKeyToDelete,

        user_id: userId,
      }),
      headers: {
        "Content-Type": "application/json",
      },
    })
      .then((res) => {
        if (res.status == 200) {
          toast.success("Successfully deleted API key.");
        } else {
          toast.error("Unable to delete an API Key.");
        }
      })
      .catch((error) => {
        toast.error("Error: ", error);
      });
  };

  // modal for delete
  const [apiKeyToDelete, setDeleteApiKey] = useState<string>("");
  const { isOpen, onOpen, onOpenChange } = useDisclosure();
  return (
    <div>
      <Toaster />
      <section className="flex p-3">
        <ListboxWrapper>
          <Listbox variant="flat" aria-label="Listbox menu with sections">
            <ListboxSection title="Actions" showDivider>
              <ListboxItem
                key="new"
                description="Create a new API Key"
                startContent={<AddNoteIcon className={iconClasses} />}
                onClick={GenerateAndInsertApiKey}
              >
                New API Key
              </ListboxItem>
            </ListboxSection>
          </Listbox>
        </ListboxWrapper>
        <Spacer x={2} className="w-full">
          {/*Delete Modal*/}
          <Modal isOpen={isOpen} onOpenChange={onOpenChange}>
            <ModalContent>
              {(onClose) => (
                <>
                  <ModalHeader className="flex flex-col gap-1">
                    Delete Api Key
                  </ModalHeader>
                  <ModalBody>
                    <p>Are you sure you want to delete? </p>
                    <span className="font-semibold text-cyan-50">
                      {apiKeyToDelete}
                    </span>
                  </ModalBody>
                  <ModalFooter>
                    <Button color="danger" variant="light" onPress={onClose}>
                      Close
                    </Button>
                    <Button
                      color="primary"
                      onPress={() => {
                        DeleteApiKey();
                        onClose();
                      }}
                    >
                      Delete
                    </Button>
                  </ModalFooter>
                </>
              )}
            </ModalContent>
          </Modal>

          <Table isStriped aria-label="Api Key Table" className="w-full">
            <TableHeader columns={columns}>
              {(column) => (
                <TableColumn
                  key={column.uid}
                  align={column.uid === "actions" ? "center" : "start"}
                >
                  {column.name}
                </TableColumn>
              )}
            </TableHeader>
            <TableBody
              items={apiKeyData}
              emptyContent={"No Api Keys have been generated yet. "}
            >
              {(item) => (
                <TableRow key={item.id}>
                  <TableCell>{item.api_key}</TableCell>
                  <TableCell>
                    <Chip
                      className="capitalize"
                      color={"success"}
                      size="sm"
                      variant="flat"
                    >
                      {"Active"}
                    </Chip>
                  </TableCell>
                  <TableCell>
                    {new Date(item.created_at).toLocaleString("en-us", {
                      dateStyle: "short",
                      timeStyle: "short",
                    })}
                  </TableCell>
                  <TableCell>
                    <div className="relative flex items-center gap-4">
                      <Tooltip content="Copy">
                        <span className="text-lg text-default-400 cursor-pointer active:opacity-50">
                          <ClipboardIcon
                            onClick={() => {
                              navigator.clipboard.writeText(item.api_key);
                            }}
                          />
                        </span>
                      </Tooltip>
                      <Tooltip color="danger" content="Delete API key">
                        <span className="text-lg text-danger cursor-pointer active:opacity-50">
                          <DeleteIcon
                            onClick={() => {
                              setDeleteApiKey(item.api_key);
                              onOpen();
                            }}
                          />
                        </span>
                      </Tooltip>
                    </div>
                  </TableCell>
                </TableRow>
              )}
            </TableBody>
          </Table>
        </Spacer>
      </section>
    </div>
  );
}

const ListboxWrapper = ({ children }: any) => (
  <div className="w-full max-w-[260px] border-small px-1 py-2 rounded-small border-default-200 dark:border-default-100">
    {children}
  </div>
);
