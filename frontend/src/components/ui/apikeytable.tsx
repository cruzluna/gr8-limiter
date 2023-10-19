"use client";
import {
  Chip,
  Input,
  Listbox,
  ListboxItem,
  ListboxSection,
  Spacer,
  Table,
  TableBody,
  TableCell,
  TableColumn,
  TableHeader,
  TableRow,
  Tooltip,
  cn,
} from "@nextui-org/react";
import { AddNoteIcon, DeleteDocumentIcon, DeleteIcon, EyeIcon } from "./icons";
import { ApiKeyPayload } from "@/app/(user)/dashboard/page";
import { UUID } from "crypto";

type ApiKeyTableProps = {
  userId: string;
  apiKeyData: ApiKeyPayload[];
};

const columns = [
  { name: "API KEY", uid: "apikey" },
  { name: "STATUS", uid: "status" },
  { name: "ACTIONS", uid: "actions" },
];

export default function ApiKeyTable({ userId, apiKeyData }: ApiKeyTableProps) {
  const iconClasses =
    "text-xl text-default-500 pointer-events-none flex-shrink-0";

  return (
    <div>
      <section className="flex p-3">
        <ListboxWrapper>
          <Listbox variant="flat" aria-label="Listbox menu with sections">
            <ListboxSection title="Actions" showDivider>
              {/* TODO: switch addnote icon to a key icon*/}

              {/*POC*/}
              <ListboxItem
                key="new"
                description="Create a new API Key"
                startContent={<AddNoteIcon className={iconClasses} />}
              >
                New API Key
              </ListboxItem>
            </ListboxSection>
            <ListboxSection title="Danger zone">
              <ListboxItem
                key="delete"
                className="text-danger"
                color="danger"
                description="Permanently delete an API Key"
                startContent={
                  <DeleteDocumentIcon
                    className={cn(iconClasses, "text-danger")}
                  />
                }
              >
                Delete API Key
              </ListboxItem>
            </ListboxSection>
          </Listbox>
        </ListboxWrapper>
        <Spacer x={2} className="w-full">
          <Table isStriped aria-label="Skeleton Table" className="w-full">
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
                    <div className="relative flex items-center gap-4">
                      <Tooltip content="Visible">
                        <span className="text-lg text-default-400 cursor-pointer active:opacity-50">
                          <EyeIcon />
                        </span>
                      </Tooltip>
                      <Tooltip color="danger" content="Delete API key">
                        <span className="text-lg text-danger cursor-pointer active:opacity-50">
                          <DeleteIcon />
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

const ListboxWrapper = ({ children }) => (
  <div className="w-full max-w-[260px] border-small px-1 py-2 rounded-small border-default-200 dark:border-default-100">
    {children}
  </div>
);

// TODO: Hides Api Key
const HiddenTextWrapper = ({
  isVisible,
  apiKey,
}: {
  isVisible: boolean;
  apiKey: UUID;
}) => {
  console.log(isVisible);
  return (
    <Input type={isVisible ? "text" : "password"} value={apiKey} isReadOnly />
  );
};
