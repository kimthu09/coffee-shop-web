// utils/withAuth.ts
import { useEffect, useState } from "react";
import { useRouter } from "next/navigation";
import { getToken } from "./auth";
import Loading from "@/components/loading";
import { useAuth } from "@/hooks/auth-context";

type WithAuthProps = {
  // Add any additional props you need to pass to the WrappedComponent
};

type ComponentType = (props: any) => JSX.Element;

const withAuth = <P extends WithAuthProps>(WrappedComponent: ComponentType) => {
  const Wrapper = (props: any) => {
    const router = useRouter();
    const [isLoading, setIsLoading] = useState(true);
    const { login } = useAuth();
    useEffect(() => {
      const token = getToken();

      // Ensure that the code below runs only on the client side
      if (typeof window !== "undefined") {
        // If the user is not authenticated, redirect to the login page
        if (!token) {
          router.push("/login");
        } else {
          login("");
          setIsLoading(false); // Set loading to false once authentication is checked
        }
      }
    }, [router]);

    return isLoading ? <Loading /> : <WrappedComponent {...props} />;
  };

  return Wrapper;
};

export default withAuth;
