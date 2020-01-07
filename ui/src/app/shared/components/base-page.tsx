import {AppContext} from 'argo-ui/src/index';
import * as PropTypes from 'prop-types';
import * as React from 'react';
import {RouteComponentProps} from 'react-router';

export class BasePage<P extends RouteComponentProps<any>, S> extends React.Component<P, S> {
    public static contextTypes = {
        router: PropTypes.object,
        apis: PropTypes.object
    };

    public getParam(name: string) {
        return new URLSearchParams(this.appContext.router.route.location.search).get(name);
    }

    // this allows us to set-multiple parameters at once
    public setParams(newParams: any) {
        const params = new URLSearchParams(this.appContext.router.route.location.search);
        Object.keys(newParams).forEach(name => {
            const value = newParams[name];
            if (value !== null) {
                params.set(name, value);
            } else {
                params.delete(name);
            }
        });
        this.appContext.router.history.push(`${this.props.match.url}?${params.toString()}`);
    }

    get appContext() {
        return this.context as AppContext;
    }
}
