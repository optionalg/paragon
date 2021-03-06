import { useQuery } from "@apollo/react-hooks";
import gql from "graphql-tag";
import * as React from "react";
import { Dropdown, DropdownProps, Input } from "semantic-ui-react";
import { Service } from "../../graphql/models";
// Suggest services for the typeahead.
export const SUGGEST_SERVICES_QUERY = gql`
  query SuggestServices {
    services {
      name
      id
      tag {
        id
        name
      }
    }
  }
`;

type ServicesResult = {
  services: Service[];
};

// XServiceTypeahead adds a service tag field to a form, which has a value of a single tag id.
const XServiceTypeahead: React.FC<{ value: string, onChange: (event: React.SyntheticEvent<HTMLElement>, data: DropdownProps) => void, labeled?: boolean }> = ({ value, onChange, labeled }) => {
  const { loading, error, data: { services = [] } = {} } = useQuery<
    ServicesResult
  >(SUGGEST_SERVICES_QUERY);

  let options = services
    ? Array.from(
      new Map(
        services.map(svc => {
          return [
            svc.tag ? svc.tag.id : null,
            {
              text: svc.name || "unknown-service",
              value: svc.tag ? svc.tag.id : null
            }
          ];
        })
      ).values()
    )
    : [];

  const getDropdown = () => (
    <Dropdown
      placeholder="Any Service"
      icon=""
      fluid
      clearable
      selection
      error={error !== null && error !== undefined}
      loading={loading}
      options={options}
      name="service_tag"
      value={value}
      onChange={onChange}
      style={{
        borderRadius: "0 4px 4px 0"
      }}
    />
  );

  if (labeled) {
    return <Input fluid label="Service" icon="cloud" input={getDropdown()} />;
  }
  return getDropdown();
};

export default XServiceTypeahead;
