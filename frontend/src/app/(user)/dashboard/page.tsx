"use client";
import { useCallback } from "react";

import {
  AddNoteIcon,
  CopyDocumentIcon,
  DeleteDocumentIcon,
  DeleteIcon,
  EditDocumentIcon,
  EditIcon,
  EyeIcon,
} from "@/components/ui/icons";

import {
  Card,
  CardHeader,
  Chip,
  ChipProps,
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
  User,
  cn,
} from "@nextui-org/react";
import { columns, users } from "./data";

const statusColorMap: Record<string, ChipProps["color"]> = {
  active: "success",
  paused: "danger",
  warning: "warning",
};
type User = (typeof users)[0];
export default function Page() {
  const iconClasses =
    "text-xl text-default-500 pointer-events-none flex-shrink-0";
  const renderCell = useCallback((user: User, columnKey: React.Key) => {
    const cellValue = user[columnKey as keyof User];
    switch (columnKey) {
      case "name":
        return (
          <User
            avatarProps={{ radius: "lg", src: user.avatar }}
            description={user.email}
            name={cellValue}
          >
            {user.email}
          </User>
        );
      case "status":
        return (
          <Chip
            className="capitalize"
            color={statusColorMap[user.status]}
            size="sm"
            variant="flat"
          >
            {cellValue}
          </Chip>
        );
      case "actions":
        return (
          <div className="relative flex items-center gap-2">
            <Tooltip content="Details">
              <span className="text-lg text-default-400 cursor-pointer active:opacity-50">
                <EyeIcon />
              </span>
            </Tooltip>
            <Tooltip content="Edit user">
              <span className="text-lg text-default-400 cursor-pointer active:opacity-50">
                <EditIcon />
              </span>
            </Tooltip>
            <Tooltip color="danger" content="Delete user">
              <span className="text-lg text-danger cursor-pointer active:opacity-50">
                <DeleteIcon />
              </span>
            </Tooltip>
          </div>
        );
      default:
        return cellValue;
    }
  }, []);

  return (
    <div>
      <section>
        <div className="flex items-center justify-center pt-[1rem]">
          <div className="flex flex-row flex-grow-0 gap-x-4 w-full p-3">
            <Card className="w-1/3 py-4" isPressable>
              <CardHeader className="pb-0 pt-2 px-4 flex-col items-start">
                <small className="text-default-500">12 Integrations</small>
                <h4 className="font-bold text-large">Manage Integrations</h4>
              </CardHeader>
            </Card>

            <Card className="w-1/3 py-4" isPressable>
              <CardHeader className="pb-0 pt-2 px-4 flex-col items-start">
                <small className="text-default-500">12 Endpoints</small>
                <h4 className="font-bold text-large">View Metrics</h4>
              </CardHeader>
            </Card>
            <Card className="w-1/3 py-4" isPressable>
              <CardHeader className="pb-0 pt-2 px-4 flex-col items-start">
                <small className="text-default-500">username</small>
                <h4 className="font-bold text-large">Manage Profile</h4>
              </CardHeader>
            </Card>
          </div>
        </div>
      </section>
      <section className="flex p-3">
        <ListboxWrapper>
          <Listbox variant="flat" aria-label="Listbox menu with sections">
            <ListboxSection title="Actions" showDivider>
              {/* TODO: switch addnote icon to a key icon*/}
              <ListboxItem
                key="new"
                description="Create a new API Key"
                startContent={<AddNoteIcon className={iconClasses} />}
              >
                New API Key
              </ListboxItem>
              <ListboxItem
                key="copy"
                description="Copy SDK generated API key"
                startContent={<CopyDocumentIcon className={iconClasses} />}
              >
                Copy Stratus API key
              </ListboxItem>
              <ListboxItem
                key="edit"
                description="Allows you to edit the file"
                startContent={<EditDocumentIcon className={iconClasses} />}
              >
                Edit API Key
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
            <TableBody items={users}>
              {(item) => (
                <TableRow key={item.id}>
                  {(columnKey) => (
                    <TableCell>{renderCell(item, columnKey)}</TableCell>
                  )}
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
