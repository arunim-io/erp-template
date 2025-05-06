import { Button } from "$components/ui/button";
import {
  Dialog,
  DialogClose,
  DialogContent,
  DialogDescription,
  DialogFooter,
  DialogHeader,
  DialogTitle,
  DialogTrigger,
} from "$components/ui/dialog";
import { graphql } from "$lib/gql";
import { authStore } from "$src/lib/stores";
import { useMutation } from "@tanstack/react-query";
import { useNavigate, useRouteContext } from "@tanstack/react-router";

const logoutMutationDocument = graphql(`
  mutation UserLogout {
    auth {
      logout
    }
  }
`);

export function LogoutDialog() {
  const navigate = useNavigate();
  const { graphqlClient, queryClient } = useRouteContext({ from: "/_app/" });

  const { mutateAsync: logoutUser } = useMutation({
    mutationKey: ["auth", "user", "logout"],
    mutationFn: () => graphqlClient.request(logoutMutationDocument),
    onSuccess() {
      queryClient.invalidateQueries({ queryKey: ["auth", "user", "status"] });
      authStore.setState(state => ({ ...state, isAuthenticated: false }));
      navigate({ to: "/login" });
    },
  });

  return (
    <Dialog>
      <DialogTrigger asChild>
        <Button>Logout</Button>
      </DialogTrigger>
      <DialogContent>
        <DialogHeader>
          <DialogTitle>Are you sure you want to sign out?</DialogTitle>
          <DialogDescription>You will need to log in again!</DialogDescription>
        </DialogHeader>
        <DialogFooter>
          <DialogClose asChild>
            <Button variant="outline">Cancel</Button>
          </DialogClose>
          <Button variant="destructive" onClick={() => logoutUser()}>Confirm</Button>
        </DialogFooter>
      </DialogContent>
    </Dialog>
  );
}
